<div align="center">
 <img alt="ollama" height="200px" src="https://github.com/ollama/ollama/assets/3325447/0d0b44e2-8f4a-4e99-9b52-a5c1c741c8f7">
</div>

# Ollama
#### Get up and running with large language models.

[![Discord](https://dcbadge.vercel.app/api/server/ollama?style=flat&compact=true)](https://discord.gg/ollama)




Ollama is a lightweight, extensible framework that enables you to effortlessly run a variety of open-source LLMs. With Ollama, you can explore the capabilities of LLMs, whether for development, research, or just for fun.

It is free and open-source and serves as a wrapper by providing a simple API for creating, running, and managing models

---

### Table of Contents

- 🛠️ [Getting Started](#getting-started)
  - 📥 [Installation](#installation)
  - 🚀 [Quickstart](#quickstart)
- 📚 [Model Library](#model-library)
- 📦 [Features](#features)
  - ⚡ [Model Execution](#model-execution)
  - 🧩 [Model Management](#model-management)
  - ✨ [Customize a Model](#customize-a-model)
  - 🎨 [Customize a Prompt](#customize-a-prompt)
  - 🖥️ [Start Ollama Service](#start-ollama-service) 
  - 🌟 [Special Cases](#special-cases)
- 🧱️ [Building](#building)
- 🔗 [REST API](#rest-api)
- 🌐 [Community Integrations](#community-integrations)
- 🔄 [Contributing](#contributing)
- 🤝 [Acknowledgments](#acknowledgments)
- 🐞 [Issues](#issues)
- 📜 [License](#license)

---

## 🛠️ Getting Started
### 📥 Installation

| Platform    | Installation Instructions                                                                                                                                                                                                                                                    |
|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **macOS**   | 🍏 [Download](https://ollama.com/download/Ollama-darwin.zip)                                                                                                                                                                                                                 |
| **Windows** | 🪟 [Download](https://ollama.com/download/OllamaSetup.exe)                                                                                                                                                                                                                   |
| **Linux**   | 🐧 Install via command line: <br/> ```bash curl -fsSL https://ollama.com/install. &#124; sh ``` <br/> <br/> 📖 [Manual install instructions](https://github.com/ollama/ollama/blob/main/docs/linux.md).                                                                           |
| **Docker**  | 📦 **Pull** the official Ollama Docker image: <br> ```bash docker pull ollama/ollama``` <br> <br/> 🐳 **Run** the Docker container: <br> ```bash docker run -it ollama/ollama ``` <br> <br/>🌐 [Explore the Ollama Docker Hub page](https://hub.docker.com/r/ollama/ollama). |


### 🚀 Quickstart

To run and chat with [Llama 3.2](https://ollama.com/library/llama3.2) :

```
ollama run llama3.2
```

🦙💥 Happy llama

## 📚 Model library

Ollama supports a list of models available on [ollama.com/library](https://ollama.com/library 'ollama model library')

Here are some example models that can be downloaded:

| Model              | Parameters | Size  | Download                       |
| ------------------ | ---------- | ----- | ------------------------------ |
| Llama 3.2          | 3B         | 2.0GB | `ollama run llama3.2`          |
| Llama 3.2          | 1B         | 1.3GB | `ollama run llama3.2:1b`       |
| Llama 3.1          | 8B         | 4.7GB | `ollama run llama3.1`          |
| Llama 3.1          | 70B        | 40GB  | `ollama run llama3.1:70b`      |
| Llama 3.1          | 405B       | 231GB | `ollama run llama3.1:405b`     |
| Phi 3 Mini         | 3.8B       | 2.3GB | `ollama run phi3`              |
| Phi 3 Medium       | 14B        | 7.9GB | `ollama run phi3:medium`       |
| Gemma 2            | 2B         | 1.6GB | `ollama run gemma2:2b`         |
| Gemma 2            | 9B         | 5.5GB | `ollama run gemma2`            |
| Gemma 2            | 27B        | 16GB  | `ollama run gemma2:27b`        |
| Mistral            | 7B         | 4.1GB | `ollama run mistral`           |
| Moondream 2        | 1.4B       | 829MB | `ollama run moondream`         |
| Neural Chat        | 7B         | 4.1GB | `ollama run neural-chat`       |
| Starling           | 7B         | 4.1GB | `ollama run starling-lm`       |
| Code Llama         | 7B         | 3.8GB | `ollama run codellama`         |
| Llama 2 Uncensored | 7B         | 3.8GB | `ollama run llama2-uncensored` |
| LLaVA              | 7B         | 4.5GB | `ollama run llava`             |
| Solar              | 10.7B      | 6.1GB | `ollama run solar`             |

> [!NOTE]
> You should have at least 8 GB of RAM available to run the 7B models, 16 GB to run the 13B models, and 32 GB to run the 33B models.


## 📦 Features 

### ⚡ Model execution

#### Launch a model

```
ollama run llama3.2
```


#### Terminate a running model

```
ollama stop llama3.2
```

---

### 🧩 Model management
#### Create a new model

`ollama create` is used to create a model from a _Modelfile_.

```
ollama create mymodel -f ./Modelfile
```

#### Pull (retrieve) a model

```
ollama pull llama3.2
```

> This command can also be used to update a local model. Only the diff will be pulled.

#### Remove a model

```
ollama rm llama3.2
```

#### Copy a model

```
ollama cp llama3.2 my-model
```

#### List all models on your system

```
ollama list
```

#### Show model information

```
ollama show llama3.2
```

#### List which models are currently loaded

```
ollama ps
```
---

### ✨ Customize a model

#### Import from GGUF

Ollama supports importing GGUF models in the Modelfile:

1. Create a file named `Modelfile`, with a `FROM` instruction with the local filepath to the model you want to import.

   ```
   FROM ./vicuna-33b.Q4_0.gguf
   ```

2. Create the model in Ollama

   ```
   ollama create example -f Modelfile
   ```

3. Run the model

   ```
   ollama run example
   ```

#### Import from PyTorch or Safetensors

See the [guide](docs/import.md) on importing models for more information.

---
### 🎨 Customize a prompt

Models from the Ollama library can be customized with a prompt. For example, to customize the `llama3.2` model:

```
ollama pull llama3.2
```

Create a `Modelfile`:

```
FROM llama3.2

# set the temperature to 1 [higher is more creative, lower is more coherent]
PARAMETER temperature 1

# set the system message
SYSTEM """
You are Mario from Super Mario Bros. Answer as Mario, the assistant, only.
"""
```

Next, create and run the model:

```
ollama create mario -f ./Modelfile
ollama run mario
>>> hi
Hello! It's your friend Mario.
```

For more examples, see the [examples](examples) directory. For more information on working with a Modelfile, see the [Modelfile](docs/modelfile.md) documentation.

---

### 🖥️ Start Ollama Service

When you want to start ollama without running the desktop application:
```
ollama serve
```
---

### 🌟 Special cases
#### 📜 Multiline input

For multiline input, you can wrap text with `"""`:

```
>>> """Hello,
... world!
... """
I'm a basic program that prints the famous "Hello, world!" message to the console.
```

#### 🖼️ Multimodal models

```
ollama run llava "What's in this image? /Users/jmorgan/Desktop/smile.png"
The image features a yellow smiley face, which is likely the central focus of the picture.
```

#### 📨 Pass the prompt as an argument

```
$ ollama run llama3.2 "Summarize this file: $(cat README.md)"
 Ollama is a lightweight, extensible framework for building and running language models on the local machine. It provides a simple API for creating, running, and managing models, as well as a library of pre-built models that can be easily used in a variety of applications.
```



## 🧱️ Building

See the [developer guide](https://github.com/ollama/ollama/blob/main/docs/development.md)

### Running local builds

Next, start the server:

```
./ollama serve
```

Finally, in a separate shell, run a model:

```
./ollama run llama3.2
```

## 🔗 REST API

Ollama has a REST API for running and managing models.

### Generate a response

```
curl http://localhost:11434/api/generate -d '{
  "model": "llama3.2",
  "prompt":"Why is the sky blue?"
}'
```

### Chat with a model

```
curl http://localhost:11434/api/chat -d '{
  "model": "llama3.2",
  "messages": [
    { "role": "user", "content": "why is the sky blue?" }
  ]
}'
```

See the [API documentation](./docs/api.md) for all endpoints.

## 🌐 Community Integrations

### **Web Applications**
- [Open WebUI](https://github.com/open-webui/open-webui)
- [Lollms-Webui](https://github.com/ParisNeo/lollms-webui)
- [Chatbot UI](https://github.com/ivanfioravanti/chatbot-ollama)
- [Chatbot UI v2](https://github.com/mckaywrigley/chatbot-ui)
- [Typescript UI](https://github.com/ollama-interface/Ollama-Gui?tab=readme-ov-file)
- [Minimalistic React UI for Ollama Models](https://github.com/richawo/minimal-llm-ui)
- [Ollama-SwiftUI](https://github.com/kghandour/Ollama-SwiftUI)
- [NextJS Web Interface for Ollama](https://github.com/jakobhoeg/nextjs-ollama-llm-ui)
- [NextChat](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web) with [Get Started Doc](https://docs.nextchat.dev/models/ollama)
- [Alpaca WebUI](https://github.com/mmo80/alpaca-webui)
- [OllamaGUI](https://github.com/enoch1118/ollamaGUI)
- [OpenAOE](https://github.com/InternLM/OpenAOE)
- [Lobe Chat](https://github.com/lobehub/lobe-chat) with [Integrating Doc](https://lobehub.com/docs/self-hosting/examples/ollama)
- [ChatOllama](https://github.com/sugarforever/chat-ollama) (Open Source Chatbot based on Ollama with Knowledge Bases)
- [CRAG Ollama Chat](https://github.com/Nagi-ovo/CRAG-Ollama-Chat) (Simple Web Search with Corrective RAG)
- [RAGFlow](https://github.com/infiniflow/ragflow) (Open-source Retrieval-Augmented Generation engine based on deep document understanding)
- [StreamDeploy](https://github.com/StreamDeploy-DevRel/streamdeploy-llm-app-scaffold) (LLM Application Scaffold)
- [chat](https://github.com/swuecho/chat) (chat web app for teams)
- [Ollama RAG Chatbot](https://github.com/datvodinh/rag-chatbot.git) (Local Chat with multiple PDFs using Ollama and RAG)
- [Archyve](https://github.com/nickthecook/archyve) (RAG-enabling document library)

### **Desktop Applications**

#### **Chat Interfaces**
- [Amica](https://github.com/semperai/amica)
- [Chatbox](https://github.com/Bin-Huang/Chatbox)
- [Dify.AI](https://github.com/langgenius/dify)
- [Hollama](https://github.com/fmaclen/hollama)
- [LibreChat](https://github.com/danny-avila/LibreChat)
- [Msty](https://msty.app)
- [Saddle](https://github.com/jikkuatwork/saddle)
- [AI Studio](https://github.com/MindWorkAI/AI-Studio)
- [LLMChat](https://github.com/trendy-design/llmchat) (Privacy focused, 100% local, intuitive all-in-one chat interface)
- [ConfiChat](https://github.com/1runeberg/confichat) (Lightweight, standalone, multi-platform, and privacy focused LLM chat interface with optional encryption)

#### **Development Tools**
- [Claude Dev](https://github.com/saoudrizwan/claude-dev) - VSCode extension for multi-file/whole-repo coding
- [QA-Pilot](https://github.com/reid41/QA-Pilot) (Chat with Code Repository)
- [Harbor](https://github.com/av/harbor) (Containerized LLM Toolkit with Ollama as default backend)
- [LLMStack](https://github.com/trypromptly/LLMStack) (No-code multi-agent framework to build LLM agents and workflows)
- [LLocal.in](https://github.com/kartikm7/llocal) (Easy to use Electron Desktop Client for Ollama)
- [crewAI with Mesop](https://github.com/rapidarchitect/ollama-crew-mesop) (Mesop Web Interface to run crewAI with Ollama)

#### **AI Tools**
- [Enchanted (macOS native)](https://github.com/AugustDev/enchanted)
- [Bionic GPT](https://github.com/bionic-gpt/bionic-gpt)
- [BrainSoup](https://www.nurgo-software.com/products/brainsoup) (Flexible native client with RAG & multi-agent automation)
- [Odin Runes](https://github.com/leonid20000/OdinRunes)
- [AnythingLLM (Docker + MacOs/Windows/Linux native app)](https://github.com/Mintplex-Labs/anything-llm)
- [Ollama Basic Chat: Uses HyperDiv Reactive UI](https://github.com/rapidarchitect/ollama_basic_chat)
- [OllamaSpring](https://github.com/CrazyNeil/OllamaSpring) (Ollama Client for macOS)
- [Sidellama](https://github.com/gyopak/sidellama) (browser-based LLM client)
- [Go-CREW](https://www.jonathanhecl.com/go-crew/) (Powerful Offline RAG in Golang)
- [Ollama4j Web UI](https://github.com/ollama4j/ollama4j-web-ui) - Java-based Web UI for Ollama built with Vaadin, Spring Boot and Ollama4j
- [PyOllaMx](https://github.com/kspviswa/pyOllaMx) - macOS application capable of chatting with both Ollama and Apple MLX models.

#### **Creative Applications**
- [Painting Droid](https://github.com/mateuszmigas/painting-droid) (Painting app with AI integrations)
- [Kerlig AI](https://www.kerlig.com/) (AI writing assistant for macOS)
- [BoltAI for Mac](https://boltai.com) (AI Chat Client for Mac)
- [Cherry Studio](https://github.com/kangfenmao/cherry-studio) (Desktop client with Ollama support)
- [PartCAD](https://github.com/openvmp/partcad/) (CAD model generation with OpenSCAD and CadQuery)
- [LLM-X](https://github.com/mrdjohnson/llm-x) (Progressive Web App)
- [Dify.AI](https://github.com/langgenius/dify)

### Terminal

- [oterm](https://github.com/ggozad/oterm)
- [Ellama Emacs client](https://github.com/s-kostyaev/ellama)
- [Emacs client](https://github.com/zweifisch/ollama)
- [gen.nvim](https://github.com/David-Kunz/gen.nvim)
- [ollama.nvim](https://github.com/nomnivore/ollama.nvim)
- [ollero.nvim](https://github.com/marco-souza/ollero.nvim)
- [ollama-chat.nvim](https://github.com/gerazov/ollama-chat.nvim)
- [ogpt.nvim](https://github.com/huynle/ogpt.nvim)
- [gptel Emacs client](https://github.com/karthink/gptel)
- [Oatmeal](https://github.com/dustinblackman/oatmeal)
- [cmdh](https://github.com/pgibler/cmdh)
- [ooo](https://github.com/npahlfer/ooo)
- [shell-pilot](https://github.com/reid41/shell-pilot)
- [tenere](https://github.com/pythops/tenere)
- [llm-ollama](https://github.com/taketwo/llm-ollama) for [Datasette's LLM CLI](https://llm.datasette.io/en/stable/).
- [typechat-cli](https://github.com/anaisbetts/typechat-cli)
- [ShellOracle](https://github.com/djcopley/ShellOracle)
- [tlm](https://github.com/yusufcanb/tlm)
- [podman-ollama](https://github.com/ericcurtin/podman-ollama)
- [gollama](https://github.com/sammcj/gollama)
- [Ollama eBook Summary](https://github.com/cognitivetech/ollama-ebook-summary/)
- [Ollama Mixture of Experts (MOE) in 50 lines of code](https://github.com/rapidarchitect/ollama_moe)
- [vim-intelligence-bridge](https://github.com/pepo-ec/vim-intelligence-bridge) Simple interaction of "Ollama" with the Vim editor

### Apple Vision Pro
- [Enchanted](https://github.com/AugustDev/enchanted)

### Database

- [MindsDB](https://github.com/mindsdb/mindsdb/blob/staging/mindsdb/integrations/handlers/ollama_handler/README.md) (Connects Ollama models with nearly 200 data platforms and apps)
- [chromem-go](https://github.com/philippgille/chromem-go/blob/v0.5.0/embed_ollama.go) with [example](https://github.com/philippgille/chromem-go/tree/v0.5.0/examples/rag-wikipedia-ollama)

### Package managers

- [Pacman](https://archlinux.org/packages/extra/x86_64/ollama/)
- [Gentoo](https://github.com/gentoo/guru/tree/master/app-misc/ollama)
- [Helm Chart](https://artifacthub.io/packages/helm/ollama-helm/ollama)
- [Guix channel](https://codeberg.org/tusharhero/ollama-guix)
- [Nix package](https://search.nixos.org/packages?channel=24.05&show=ollama&from=0&size=50&sort=relevance&type=packages&query=ollama)
- [Flox](https://flox.dev/blog/ollama-part-one)

### Libraries

- [LangChain](https://python.langchain.com/docs/integrations/llms/ollama) and [LangChain.js](https://js.langchain.com/docs/modules/model_io/models/llms/integrations/ollama) with [example](https://js.langchain.com/docs/use_cases/question_answering/local_retrieval_qa)
- [Firebase Genkit](https://firebase.google.com/docs/genkit/plugins/ollama)
- [crewAI](https://github.com/crewAIInc/crewAI)
- [LangChainGo](https://github.com/tmc/langchaingo/) with [example](https://github.com/tmc/langchaingo/tree/main/examples/ollama-completion-example)
- [LangChain4j](https://github.com/langchain4j/langchain4j) with [example](https://github.com/langchain4j/langchain4j-examples/tree/main/ollama-examples/src/main/java)
- [LangChainRust](https://github.com/Abraxas-365/langchain-rust) with [example](https://github.com/Abraxas-365/langchain-rust/blob/main/examples/llm_ollama.rs)
- [LlamaIndex](https://docs.llamaindex.ai/en/stable/examples/llm/ollama/) and [LlamaIndexTS](https://ts.llamaindex.ai/modules/llms/available_llms/ollama)
- [LiteLLM](https://github.com/BerriAI/litellm)
- [OllamaFarm for Go](https://github.com/presbrey/ollamafarm)
- [OllamaSharp for .NET](https://github.com/awaescher/OllamaSharp)
- [Ollama for Ruby](https://github.com/gbaptista/ollama-ai)
- [Ollama-rs for Rust](https://github.com/pepperoni21/ollama-rs)
- [Ollama-hpp for C++](https://github.com/jmont-dev/ollama-hpp)
- [Ollama4j for Java](https://github.com/ollama4j/ollama4j)
- [ModelFusion Typescript Library](https://modelfusion.dev/integration/model-provider/ollama)
- [OllamaKit for Swift](https://github.com/kevinhermawan/OllamaKit)
- [Ollama for Dart](https://github.com/breitburg/dart-ollama)
- [Ollama for Laravel](https://github.com/cloudstudio/ollama-laravel)
- [LangChainDart](https://github.com/davidmigloz/langchain_dart)
- [Semantic Kernel - Python](https://github.com/microsoft/semantic-kernel/tree/main/python/semantic_kernel/connectors/ai/ollama)
- [Haystack](https://github.com/deepset-ai/haystack-integrations/blob/main/integrations/ollama.md)
- [Elixir LangChain](https://github.com/brainlid/langchain)
- [Ollama for R - rollama](https://github.com/JBGruber/rollama)
- [Ollama for R - ollama-r](https://github.com/hauselin/ollama-r)
- [Ollama-ex for Elixir](https://github.com/lebrunel/ollama-ex)
- [Ollama Connector for SAP ABAP](https://github.com/b-tocs/abap_btocs_ollama)
- [Testcontainers](https://testcontainers.com/modules/ollama/)
- [Portkey](https://portkey.ai/docs/welcome/integration-guides/ollama)
- [PromptingTools.jl](https://github.com/svilupp/PromptingTools.jl) with an [example](https://svilupp.github.io/PromptingTools.jl/dev/examples/working_with_ollama)
- [LlamaScript](https://github.com/Project-Llama/llamascript)
- [Gollm](https://docs.gollm.co/examples/ollama-example)
- [Ollamaclient for Golang](https://github.com/xyproto/ollamaclient)
- [High-level function abstraction in Go](https://gitlab.com/tozd/go/fun)
- [Ollama PHP](https://github.com/ArdaGnsrn/ollama-php)
- [Agents-Flex for Java](https://github.com/agents-flex/agents-flex) with [example](https://github.com/agents-flex/agents-flex/tree/main/agents-flex-llm/agents-flex-llm-ollama/src/test/java/com/agentsflex/llm/ollama)

### Mobile

- [Enchanted](https://github.com/AugustDev/enchanted)
- [Maid](https://github.com/Mobile-Artificial-Intelligence/maid)
- [ConfiChat](https://github.com/1runeberg/confichat) (Lightweight, standalone, multi-platform, and privacy focused LLM chat interface with optional encryption)

### Extensions & Plugins

- [Raycast extension](https://github.com/MassimilianoPasquini97/raycast_ollama)
- [Discollama](https://github.com/mxyng/discollama) (Discord bot inside the Ollama discord channel)
- [Continue](https://github.com/continuedev/continue)
- [Obsidian Ollama plugin](https://github.com/hinterdupfinger/obsidian-ollama)
- [Logseq Ollama plugin](https://github.com/omagdy7/ollama-logseq)
- [NotesOllama](https://github.com/andersrex/notesollama) (Apple Notes Ollama plugin)
- [Dagger Chatbot](https://github.com/samalba/dagger-chatbot)
- [Discord AI Bot](https://github.com/mekb-turtle/discord-ai-bot)
- [Ollama Telegram Bot](https://github.com/ruecat/ollama-telegram)
- [Hass Ollama Conversation](https://github.com/ej52/hass-ollama-conversation)
- [Rivet plugin](https://github.com/abrenneke/rivet-plugin-ollama)
- [Obsidian BMO Chatbot plugin](https://github.com/longy2k/obsidian-bmo-chatbot)
- [Cliobot](https://github.com/herval/cliobot) (Telegram bot with Ollama support)
- [Copilot for Obsidian plugin](https://github.com/logancyang/obsidian-copilot)
- [Obsidian Local GPT plugin](https://github.com/pfrankov/obsidian-local-gpt)
- [Open Interpreter](https://docs.openinterpreter.com/language-model-setup/local-models/ollama)
- [Llama Coder](https://github.com/ex3ndr/llama-coder) (Copilot alternative using Ollama)
- [Ollama Copilot](https://github.com/bernardo-bruning/ollama-copilot) (Proxy that allows you to use ollama as a copilot like Github copilot)
- [twinny](https://github.com/rjmacarthy/twinny) (Copilot and Copilot chat alternative using Ollama)
- [Wingman-AI](https://github.com/RussellCanfield/wingman-ai) (Copilot code and chat alternative using Ollama and Hugging Face)
- [Page Assist](https://github.com/n4ze3m/page-assist) (Chrome Extension)
- [Plasmoid Ollama Control](https://github.com/imoize/plasmoid-ollamacontrol) (KDE Plasma extension that allows you to quickly manage/control Ollama model)
- [AI Telegram Bot](https://github.com/tusharhero/aitelegrambot) (Telegram bot using Ollama in backend)
- [AI ST Completion](https://github.com/yaroslavyaroslav/OpenAI-sublime-text) (Sublime Text 4 AI assistant plugin with Ollama support)
- [Discord-Ollama Chat Bot](https://github.com/kevinthedang/discord-ollama) (Generalized TypeScript Discord Bot w/ Tuning Documentation)
- [Discord AI chat/moderation bot](https://github.com/rapmd73/Companion) Chat/moderation bot written in python. Uses Ollama to create personalities.
- [Headless Ollama](https://github.com/nischalj10/headless-ollama) (Scripts to automatically install ollama client & models on any OS for apps that depends on ollama server)
- [vnc-lm](https://github.com/jk011ru/vnc-lm) (A containerized Discord bot with support for attachments and web links)
- [LSP-AI](https://github.com/SilasMarvin/lsp-ai) (Open-source language server for AI-powered functionality)
- [QodeAssist](https://github.com/Palm1r/QodeAssist) (AI-powered coding assistant plugin for Qt Creator)
- [Obsidian Quiz Generator plugin](https://github.com/ECuiDev/obsidian-quiz-generator)



---

<div style="max-width: 800px; margin: auto; text-align: left;">
    <h2>Community Integrations</h2>
    <div style="display: flex; flex-wrap: wrap; gap: 20px;">
        <div style="flex: 1; min-width: 250px;">
            <h3>Chat Interfaces</h3>
            <ul>
                <li><a href="https://github.com/semperai/amica">Amica</a></li>
                <li><a href="https://github.com/Bin-Huang/Chatbox">Chatbox</a></li>
                <li><a href="https://github.com/langgenius/dify">Dify.AI</a></li>
                <li><a href="https://github.com/fmaclen/hollama">Hollama</a></li>
                <li><a href="https://github.com/danny-avila/LibreChat">LibreChat</a></li>
                <li><a href="https://msty.app">Msty</a></li>
                <li><a href="https://github.com/jikkuatwork/saddle">Saddle</a></li>
                <li><a href="https://github.com/MindWorkAI/AI-Studio">AI Studio</a></li>
                <li><a href="https://github.com/trendy-design/llmchat">LLMChat</a> (Privacy focused, 100% local, intuitive all-in-one chat interface)</li>
                <li><a href="https://github.com/1runeberg/confichat">ConfiChat</a> (Lightweight, standalone, multi-platform, and privacy focused LLM chat interface with optional encryption)</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Development Tools</h3>
            <ul>
                <li><a href="https://github.com/saoudrizwan/claude-dev">Claude Dev</a> - VSCode extension for multi-file/whole-repo coding</li>
                <li><a href="https://github.com/reid41/QA-Pilot">QA-Pilot</a> (Chat with Code Repository)</li>
                <li><a href="https://github.com/av/harbor">Harbor</a> (Containerized LLM Toolkit with Ollama as default backend)</li>
                <li><a href="https://github.com/trypromptly/LLMStack">LLMStack</a> (No-code multi-agent framework to build LLM agents and workflows)</li>
                <li><a href="https://github.com/kartikm7/llocal">LLocal.in</a> (Easy to use Electron Desktop Client for Ollama)</li>
                <li><a href="https://github.com/rapidarchitect/ollama-crew-mesop">crewAI with Mesop</a> (Mesop Web Interface to run crewAI with Ollama)</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Clients</h3>
            <ul>
                <li><a href="https://github.com/AugustDev/enchanted">Enchanted (macOS native)</a></li>
                <li><a href="https://github.com/bionic-gpt/bionic-gpt">Bionic GPT</a></li>
                <li><a href="https://www.nurgo-software.com/products/brainsoup">BrainSoup</a> (Flexible native client with RAG & multi-agent automation)</li>
                <li><a href="https://github.com/leonid20000/OdinRunes">Odin Runes</a></li>
                <li><a href="https://github.com/Mintplex-Labs/anything-llm">AnythingLLM (Docker + MacOs/Windows/Linux native app)</a></li>
                <li><a href="https://github.com/rapidarchitect/ollama_basic_chat">Ollama Basic Chat</a> (Uses HyperDiv Reactive UI)</li>
                <li><a href="https://github.com/CrazyNeil/OllamaSpring">OllamaSpring</a> (Ollama Client for macOS)</li>
                <li><a href="https://github.com/gyopak/sidellama">Sidellama</a> (browser-based LLM client)</li>
                <li><a href="https://www.jonathanhecl.com/go-crew/">Go-CREW</a> (Powerful Offline RAG in Golang)</li>
                <li><a href="https://github.com/ollama4j/ollama4j-web-ui">Ollama4j Web UI</a> - Java-based Web UI for Ollama built with Vaadin, Spring Boot and Ollama4j</li>
                <li><a href="https://github.com/kspviswa/pyOllaMx">PyOllaMx</a> - macOS application capable of chatting with both Ollama and Apple MLX models.</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Creative Applications</h3>
            <ul>
                <li><a href="https://github.com/mateuszmigas/painting-droid">Painting Droid</a> (Painting app with AI integrations)</li>
                <li><a href="https://www.kerlig.com/">Kerlig AI</a> (AI writing assistant for macOS)</li>
                <li><a href="https://boltai.com">BoltAI for Mac</a> (AI Chat Client for Mac)</li>
                <li><a href="https://github.com/kangfenmao/cherry-studio">Cherry Studio</a> (Desktop client with Ollama support)</li>
                <li><a href="https://github.com/openvmp/partcad/">PartCAD</a> (CAD model generation with OpenSCAD and CadQuery)</li>
                <li><a href="https://github.com/mrdjohnson/llm-x">LLM-X</a> (Progressive Web App)</li>
                <li><a href="https://github.com/langgenius/dify">Dify.AI</a></li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Web Applications</h3>
            <ul>
                <li><a href="https://github.com/open-webui/open-webui">Open WebUI</a></li>
                <li><a href="https://github.com/ParisNeo/lollms-webui">Lollms-Webui</a></li>
                <li><a href="https://github.com/ivanfioravanti/chatbot-ollama">Chatbot UI</a></li>
                <li><a href="https://github.com/mckaywrigley/chatbot-ui">Chatbot UI v2</a></li>
                <li><a href="https://github.com/ollama-interface/Ollama-Gui?tab=readme-ov-file">Typescript UI</a></li>
                <li><a href="https://github.com/richawo/minimal-llm-ui">Minimalistic React UI for Ollama Models</a></li>
                <li><a href="https://github.com/kghandour/Ollama-SwiftUI">Ollama-SwiftUI</a></li>
                <li><a href="https://github.com/jakobhoeg/nextjs-ollama-llm-ui">NextJS Web Interface for Ollama</a></li>
                <li><a href="https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web">NextChat</a> with <a href="https://docs.nextchat.dev/models/ollama">Get Started Doc</a></li>
                <li><a href="https://github.com/mmo80/alpaca-webui">Alpaca WebUI</a></li>
                <li><a href="https://github.com/enoch1118/ollamaGUI">OllamaGUI</a></li>
                <li><a href="https://github.com/InternLM/OpenAOE">OpenAOE</a></li>
                <li><a href="https://github.com/lobehub/lobe-chat">Lobe Chat</a> with <a href="https://lobehub.com/docs/self-hosting/examples/ollama">Integrating Doc</a></li>
                <li><a href="https://github.com/sugarforever/chat-ollama">ChatOllama</a> (Open Source Chatbot based on Ollama with Knowledge Bases)</li>
                <li><a href="https://github.com/Nagi-ovo/CRAG-Ollama-Chat">CRAG Ollama Chat</a> (Simple Web Search with Corrective RAG)</li>
                <li><a href="https://github.com/infiniflow/ragflow">RAGFlow</a> (Open-source Retrieval-Augmented Generation engine based on deep document understanding)</li>
                <li><a href="https://github.com/StreamDeploy-DevRel/streamdeploy-llm-app-scaffold">StreamDeploy</a> (LLM Application Scaffold)</li>
                <li><a href="https://github.com/swuecho/chat">chat</a> (chat web app for teams)</li>
                <li><a href="https://github.com/datvodinh/rag-chatbot.git">Ollama RAG Chatbot</a> (Local Chat with multiple PDFs using Ollama and RAG)</li>
                <li><a href="https://github.com/nickthecook/archyve">Archyve</a> (RAG-enabling document library)</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Terminal Applications</h3>
            <ul>
                <li><a href="https://github.com/ggozad/oterm">oterm</a></li>
                <li><a href="https://github.com/s-kostyaev/ellama">Ellama Emacs client</a></li>
                <li><a href="https://github.com/zweifisch/ollama">Emacs client</a></li>
                <li><a href="https://github.com/David-Kunz/gen.nvim">gen.nvim</a></li>
                <li><a href="https://github.com/nomnivore/ollama.nvim">ollama.nvim</a></li>
                <li><a href="https://github.com/marco-souza/ollero.nvim">ollero.nvim</a></li>
                <li><a href="https://github.com/gerazov/ollama-chat.nvim">ollama-chat.nvim</a></li>
                <li><a href="https://github.com/huynle/ogpt.nvim">ogpt.nvim</a></li>
                <li><a href="https://github.com/karthink/gptel">gptel Emacs client</a></li>
                <li><a href="https://github.com/dustinblackman/oatmeal">Oatmeal</a></li>
                <li><a href="https://github.com/pgibler/cmdh">cmdh</a></li>
                <li><a href="https://github.com/npahlfer/ooo">ooo</a></li>
                <li><a href="https://github.com/reid41/shell-pilot">shell-pilot</a></li>
                <li><a href="https://github.com/pythops/tenere">tenere</a></li>
                <li><a href="https://github.com/taketwo/llm-ollama">llm-ollama</a> for <a href="https://llm.datasette.io/en/stable/">Datasette's LLM CLI</a></li>
                <li><a href="https://github.com/anaisbetts/typechat-cli">typechat-cli</a></li>
                <li><a href="https://github.com/djcopley/ShellOracle">ShellOracle</a></li>
                <li><a href="https://github.com/yusufcanb/tlm">tlm</a></li>
                <li><a href="https://github.com/ericcurtin/podman-ollama">podman-ollama</a></li>
                <li><a href="https://github.com/sammcj/gollama">gollama</a></li>
                <li><a href="https://github.com/cognitivetech/ollama-ebook-summary/">Ollama eBook Summary</a></li>
                <li><a href="https://github.com/rapidarchitect/ollama_moe">Ollama Mixture of Experts (MOE) in 50 lines of code</a></li>
                <li><a href="https://github.com/pepo-ec/vim-intelligence-bridge">vim-intelligence-bridge</a> Simple interaction of "Ollama" with the Vim editor</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Apple Vision Pro</h3>
            <ul>
                <li><a href="https://github.com/AugustDev/enchanted">Enchanted</a>: A native application for Apple Vision Pro.</li>
                <li><a href="https://github.com/mmo80/alpaca-webui">Alpaca WebUI</a>: A web-based user interface for Ollama models.</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Database Integrations</h3>
            <ul>
                <li><a href="https://github.com/mindsdb/mindsdb/blob/staging/mindsdb/integrations/handlers/ollama_handler/README.md">MindsDB</a>: Connects Ollama models with various data platforms.</li>
                <li><a href="https://github.com/philippgille/chromem-go/blob/v0.5.0/embed_ollama.go">chromem-go</a>: Integration for embedding Ollama in applications.</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Package Managers</h3>
            <ul>
                <li><a href="https://archlinux.org/packages/extra/x86_64/ollama/">Pacman</a>: Ollama package for Arch Linux.</li>
                <li><a href="https://github.com/gentoo/guru/tree/master/app-misc/ollama">Gentoo</a>: Ollama package for Gentoo Linux.</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Libraries</h3>
            <ul>
                <li><a href="https://python.langchain.com/docs/integrations/llms/ollama">LangChain</a>: Integrates Ollama with various programming languages.</li>
                <li><a href="https://github.com/tmc/langchaingo/">LangChainGo</a>: Go library for integrating with Ollama.</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Mobile Applications</h3>
            <ul>
                <li><a href="https://github.com/AugustDev/enchanted">Enchanted</a>: Mobile application for interacting with Ollama.</li>
                <li><a href="https://github.com/Mobile-Artificial-Intelligence/maid">Maid</a>: A mobile assistant application using Ollama.</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Extensions & Plugins</h3>
            <ul>
                <li><a href="https://github.com/MassimilianoPasquini97/raycast_ollama">Raycast extension</a>: Integration with Raycast for using Ollama.</li>
                <li><a href="https://github.com/hinterdupfinger/obsidian-ollama">Obsidian Ollama plugin</a>: Plugin for integrating Ollama into Obsidian.</li>
            </ul>
        </div>
        <div style="flex: 1; min-width: 250px;">
            <h3>Supported Backends</h3>
            <ul>
                <li><a href="https://github.com/ggerganov/llama.cpp">llama.cpp</a>: Project for implementing Ollama models.</li>
                <li><a href="https://github.com/rapidarchitect/ollama">Ollama API</a>: Backend API for integrating with Ollama.</li>
            </ul>
        </div>
    </div>
</div>



<div style="max-width: 800px; margin: auto; text-align: left;">
    <h2>Terminal Applications</h2>
    <div style="display: flex; flex-wrap: wrap; gap: 10px;">
        <div style="flex: 1 1 45%; min-width: 200px; border: 1px solid #ccc; padding: 10px;">
            <ul>
                <li><a href="https://github.com/ggozad/oterm">oterm</a></li>
                <li><a href="https://github.com/s-kostyaev/ellama">Ellama Emacs client</a></li>
                <li><a href="https://github.com/zweifisch/ollama">Emacs client</a></li>
                <li><a href="https://github.com/David-Kunz/gen.nvim">gen.nvim</a></li>
                <li><a href="https://github.com/nomnivore/ollama.nvim">ollama.nvim</a></li>
            </ul>
        </div>
        <div style="flex: 1 1 45%; min-width: 200px; border: 1px solid #ccc; padding: 10px;">
            <ul>
                <li><a href="https://github.com/marco-souza/ollero.nvim">ollero.nvim</a></li>
                <li><a href="https://github.com/gerazov/ollama-chat.nvim">ollama-chat.nvim</a></li>
            </ul>
        </div>
        <div style="flex: 1 1 45%; min-width: 200px; border: 1px solid #ccc; padding: 10px;">
            <ul>
                <li><a href="https://github.com/pgibler/cmdh">cmdh</a></li>
                <li><a href="https://github.com/npahlfer/ooo">ooo</a></li>
                <li><a href="https://github.com/reid41/shell-pilot">shell-pilot</a></li>
                <li><a href="https://github.com/pythops/tenere">tenere</a></li>
                <li><a href="https://github.com/taketwo/llm-ollama">llm-ollama</a> for <a href="https://llm.datasette.io/en/stable/">Datasette's LLM CLI</a></li>
            </ul>
        </div>
        <div style="flex: 1 1 45%; min-width: 200px; border: 1px solid #ccc; padding: 10px;">
            <ul>
                <li><a href="https://github.com/anaisbetts/typechat-cli">typechat-cli</a></li>
                <li><a href="https://github.com/djcopley/ShellOracle">ShellOracle</a></li>
                <li><a href="https://github.com/yusufcanb/tlm">tlm</a></li>
                <li><a href="https://github.com/ericcurtin/podman-ollama">podman-ollama</a></li>
                <li><a href="https://github.com/sammcj/gollama">gollama</a></li>
            </ul>
        </div>
        <div style="flex: 1 1 45%; min-width: 200px; border: 1px solid #ccc; padding: 10px;">
            <ul>
                <li><a href="https://github.com/cognitivetech/ollama-ebook-summary/">Ollama eBook Summary</a></li>
                <li><a href="https://github.com/rapidarchitect/ollama_moe">Ollama Mixture of Experts (MOE) in 50 lines of code</a></li>
                <li><a href="https://github.com/pepo-ec/vim-intelligence-bridge">vim-intelligence-bridge</a> Simple interaction of "Ollama" with the Vim editor</li>
            </ul>
        </div>
    </div>
</div>

### 🔄 Contributing
We welcome contributions from the community! If you’d like to help improve this project, please refer to our [CONTRIBUTING.md](CONTRIBUTING.md) document.

### 🤝 Acknowledgments

A big thank you to all the contributors who have supported this project. Your feedback, suggestions, and contributions are greatly appreciated!

### 🐞 Issues

If you encounter any issues while using the project, please follow these steps:

<div style="max-width: 700px;">
  <table style="width: 100%; border-spacing: 20px; border-collapse: separate;">
    <tr>
      <td style="text-align: center; padding: 20px; vertical-align: middle;">
        <strong>1. Check Existing Issues</strong><br>
        Start by checking the link below to see if your problem has already been reported:<br>
        <a href="https://github.com/ollama/ollama/issues" style="display: inline-block; padding: 10px 16px; margin: 16px; background-color: #007BFF; color: white; text-decoration: none; border-radius: 5px; font-size: 14px; height: 40px; line-height: 20px;">🔍 View Existing Issues</a>
      </td>
      <td style="text-align: center; padding: 20px; vertical-align: middle;">
        <strong>2. Report a New Issue</strong><br>
        If you don't find a match, please open a new issue using the link below:<br>
        <a href="https://github.com/ollama/ollama/issues/new" style="display: inline-block; padding: 10px 16px; margin: 16px; background-color: #FF0000; color: white; text-decoration: none; border-radius: 5px; font-size: 14px; height: 40px; line-height: 20px;">📝 Open a New Issue</a>
      </td>
    </tr>
  </table>
</div>

Your feedback is valuable and helps us improve the project—thank you!


## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) for details.
