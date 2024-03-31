# Instructions to use Chunking Tool

# Overview
CHUNK-APP is a specialized tool designed to manage and process large datasets. It divides datasets into smaller, more manageable chunks, making them easier to handle and work with. This division not only aids in data management but also updates metadata to accurately reflect these subdivisions. Each data chunk receives its own metadata file, ensuring comprehensive documentation and independent usability. This is part of a larger tool set provided by Merkle META.

#Structure
The application is organized into several directories and files, each serving a specific purpose in the dataset processing pipeline:

- CHUNK_APP/cmd/main.go: The entry point of the application, orchestrating the data chunking process.
- CHUNK_APP/pkg/parsechunk/parsechunk.go: Contains the logic for parsing and chunking the dataset, including metadata updates.
- CHUNK_APP/data: Directory for input data files. This includes both the dataset to be chunked and the associated metadata in JSON format.
- CHUNK_APP/jsondata/: This directory holds the JSON data produced after chunking. It's populated by the combinedjsoncreation.js script, which combines chunked data into a - single JSON file.
- CHUNK_APP/dataout/: Output directory for parsed and chunked data. This is where the final processed files are stored.

## Features

CHUNK-APP offers a robust suite of features designed to efficiently manage and process large datasets:

- **Data Chunking**: At its core, CHUNK_APP excels in breaking down large datasets into smaller, more manageable chunks. This process is vital for improving the accessibility and usability of big data, especially in environments with limited processing capabilities.

- **Metadata Management**: Alongside data chunking, the application automatically updates and manages metadata for each chunk. This ensures every piece of data is adequately described, facilitating easier access, analysis, and independent use.

- **Automatic JSON Handling**: Utilizing a JavaScript utility (`combinedjsoncreation.js`), CHUNK_APP automates the process of generating combined JSON files for chunked data. This feature streamlines the workflow for users, making data integration and further processing seamless.

- **Flexible Data Input and Output**: With designated directories for input data (`data`), intermediate JSON data (`jsondata`), and output chunked data (`dataout`), the application maintains a clear and organized data processing pipeline. This organization supports efficient data management and accessibility.

- **Simplified Setup and Operation**: Designed with user-friendliness in mind, CHUNK_APP offers a straightforward setup process, requiring only Go and Git as prerequisites. The application can be quickly deployed and run, making it accessible even to those with limited programming experience.

- **Customizable Data Processing**: While CHUNK_APP comes pre-configured for general use, its Go and JavaScript components can be easily modified to meet specific data processing needs. This flexibility allows users to tailor the application to their unique dataset requirements.

- **Open Source Collaboration**: As an open-source project, CHUNK_APP encourages community collaboration. Users can contribute to its development by suggesting new features, improvements, or fixes, fostering a collaborative environment for enhancing the tool's capabilities.

## Setup

### Prerequisites

- Go (version 1.15 or newer)
- Git

### Installation

1. **Clone the repository**:

```bash
git clone {https://github.com/ShaneSCalder/MAC-ToolBox/edit/main/ToolBox/ChunkingTool/CHUNK_APP}
```

2. **Build the application** (if applicable):

```bash
cd CHUNK-APP
```

```bash
go mod init CHUNK-APP
```

```bash
go mod tidy
```
3. **Run the server**:

```bash
go run cmd/main.go
```

## Next Steps

As CHUNK_APP continues to evolve, we are focusing on enhancements and new features that will extend its utility across various domains, including enterprise solutions and tools for content creators and data labelers. Here are some of the key areas we're working on:

### 1. Scalability Improvements

To accommodate even larger datasets, we are exploring advanced data processing algorithms and optimization techniques. Our goal is to increase the efficiency and speed of data chunking and metadata management, ensuring the tool can handle multiple teraflops of data with ease.

### 2. Enhanced User Interface

While CHUNK_APP is currently a command-line tool, we plan to develop a user-friendly graphical interface. This interface will make it easier for users to manage their data processing tasks, configure settings, and monitor the progress of data chunking operations.

### 3. Integration Capabilities

We aim to introduce API access and integration capabilities with popular data storage and processing platforms. This will enable users to seamlessly incorporate CHUNK-APP into their existing data analysis and management workflows.

### 4. Expanded Support for Data Formats

Recognizing the diverse nature of datasets, we are working to support additional data formats, especially those commonly used in AI training, such as images and videos. This expansion will make CHUNK-APP a more versatile tool for content creators and data labelers.

### 5. Community and Enterprise Editions

To cater to different user needs, we plan to offer both Community and Enterprise editions of CHUNK_APP. The Enterprise edition will include advanced features, dedicated support, and customization options suitable for large-scale, commercial environments.

### 6. Open Source Collaboration

We believe in the power of community-driven development. To this end, we will continue to foster an open-source culture, encouraging contributions, feedback, and ideas from users and developers. We look forward to collaborating on new features, bug fixes, and performance enhancements.

### 7. Documentation and Tutorials

To assist users in maximizing the tool's potential, we will expand our documentation, offering comprehensive guides, usage scenarios, and tutorials. This resource expansion will help new users get started and enable advanced users to leverage CHUNK-APP's full capabilities.

---

Your feedback and contributions are invaluable to us as we take these next steps. Together, we can build a powerful toolbox that meets the growing needs of data processing and analysis across various industries.

# Work Flow 

![MAC_POC_MerkleMETA](https://github.com/ShaneSCalder/MAC-ToolBox/assets/29208274/f2691e30-db14-4c84-a2fc-3b1e87cdb7e0)

# Use case Data Sale 

![ChunkedDataSale_UseCase_MerkleMETA](https://github.com/ShaneSCalder/MAC-ToolBox/assets/29208274/9c95e558-e56c-47ad-ac45-daef630196bf)

