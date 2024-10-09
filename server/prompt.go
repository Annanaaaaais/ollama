package server

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ollama/ollama/api"
	"github.com/ollama/ollama/llm"
	"github.com/ollama/ollama/server/imageproc"
	"github.com/ollama/ollama/template"
)

type tokenizeFunc func(context.Context, string) ([]int, error)

// chatPrompt accepts a list of messages and returns the prompt and images that should be used for the next chat turn.
// chatPrompt truncates any messages that exceed the context window of the model, making sure to always include 1) the
// latest message and 2) system messages
func chatPrompt(ctx context.Context, m *Model, tokenize tokenizeFunc, opts *api.Options, msgs []api.Message, tools []api.Tool) (prompt string, images []llm.ImageData, _ error) {
	var system []api.Message

	// always include the last message
	n := len(msgs) - 1
	// in reverse, find all messages that fit into context window
	for i := n - 1; i >= 0; i-- {
		system = make([]api.Message, 0)
		for j := range i {
			if msgs[j].Role == "system" {
				system = append(system, msgs[j])
			}
		}

		var b bytes.Buffer
		if err := m.Template.Execute(&b, template.Values{Messages: append(system, msgs[i:]...), Tools: tools}); err != nil {
			return "", nil, err
		}

		s, err := tokenize(ctx, b.String())
		if err != nil {
			return "", nil, err
		}

		ctxLen := len(s)
		if m.ProjectorPaths != nil {
			for _, m := range msgs[i:] {
				// images are represented as 768 sized embeddings
				// TODO: get embedding length from project metadata
				ctxLen += 768 * len(m.Images)
			}
		}

		if ctxLen > opts.NumCtx {
			slog.Debug("truncating input messages which exceed context length", "truncated", len(msgs[i:]))
			break
		} else {
			n = i
		}
	}

	currMsgIdx := n

	if checkMllamaModelFamily(m) {
		// Mllama only supports a single image so include the first one of the most recent message
		lastMsgIdx := len(msgs) - 1
		for cnt := lastMsgIdx; cnt >= currMsgIdx; cnt-- {
			msg := msgs[cnt]
			if len(msg.Images) >= 1 {
				data, aspectRatioID, err := imageproc.Preprocess(msg.Images[0])
				if err != nil {
					return "", nil, err
				}

				buf := new(bytes.Buffer)
				err = binary.Write(buf, binary.LittleEndian, data)
				if err != nil {
					return "", nil, err
				}

				imgData := llm.ImageData{
					Data:          buf.Bytes(),
					AspectRatioID: aspectRatioID,
				}

				msgs[cnt].Content = strings.TrimSpace("<|image|>" + msg.Content)
				images = append(images, imgData)

				break
			}
		}
	} else {
		for cnt, msg := range msgs[currMsgIdx:] {
			for _, i := range msg.Images {
				imgData := llm.ImageData{
					ID:   len(images),
					Data: i,
				}

				imageTag := fmt.Sprintf("[img-%d]", imgData.ID)
				prompt := msg.Content

				if !strings.Contains(prompt, "[img]") {
					prompt = strings.TrimSpace("[img] " + prompt)
				}
				prompt = strings.Replace(prompt, "[img]", imageTag, 1)
				msgs[currMsgIdx+cnt].Content = prompt

				images = append(images, imgData)
			}
		}
	}

	// truncate any messages that do not fit into the context window
	var b bytes.Buffer
	if err := m.Template.Execute(&b, template.Values{Messages: append(system, msgs[currMsgIdx:]...), Tools: tools}); err != nil {
		return "", nil, err
	}

	return b.String(), images, nil
}

func checkMllamaModelFamily(m *Model) bool {
	for _, arch := range m.Config.ModelFamilies {
		if arch == "mllama" {
			return true
		}
	}
	return false
}
