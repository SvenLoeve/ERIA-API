# Project Structure

Our project structure is a simplified version of Layered Architecture. For this project, we use the following structure:

````
.
├── main.go			    # Main go executable file
├── README.md			# Install guide
├── build-pipeline.yml	# Azure DevOps CD-pipeline
├── test-pipeline.yml	# Azure DevOps CI-pipeline
├── example.env			# Default environment variables file
├── api				    # All the API routes/handlers
├── database
│   └── ent
│        └── schema		# Object relational mapping schema
├── docs				# Writerside API documentation
├── services			# Service layer for API routes/handlers
│   ├── cryptography
│   ├── jwt
│   ├── math
│   ├── nfc
│   ├── pagination
│   └── validate
├── static        		# Static content served through HTTP server
├── tests				# All integration tests
└── types				# All structs/data-transfer-objects
````