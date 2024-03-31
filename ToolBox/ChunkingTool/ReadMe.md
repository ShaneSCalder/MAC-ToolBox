# Instructions to use Chunking Tool

# Overview
CHUNK-APP is a specialized tool designed to manage and process large datasets. It divides datasets into smaller, more manageable chunks, making them easier to handle and work with. This division not only aids in data management but also updates metadata to accurately reflect these subdivisions. Each data chunk receives its own metadata file, ensuring comprehensive documentation and independent usability.

#Structure
The application is organized into several directories and files, each serving a specific purpose in the dataset processing pipeline:

- CHUNK_APP/cmd/main.go: The entry point of the application, orchestrating the data chunking process.
- CHUNK_APP/pkg/parsechunk/parsechunk.go: Contains the logic for parsing and chunking the dataset, including metadata updates.
- CHUNK_APP/data: Directory for input data files. This includes both the dataset to be chunked and the associated metadata in JSON format.
- CHUNK_APP/jsondata/: This directory holds the JSON data produced after chunking. It's populated by the combinedjsoncreation.js script, which combines chunked data into a - single JSON file.
- CHUNK_APP/dataout/: Output directory for parsed and chunked data. This is where the final processed files are stored.
