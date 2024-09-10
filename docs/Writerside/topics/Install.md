# Install
There are two ways to install the backend. We recommend using Docker, as it is faster and easier than the alternative method.
##### Docker (Recommended)
1. Rename ```example.env``` to ```.env```
2. Run ``` docker compose build```
3. Next, run the project with ```docker compose up ```

##### Local
1. Rename ```example.env``` to ```.env```
2. Setup environment variables inside the ```.env``` file
   - Specify the database port and IP address to be used by the API.
   - Turn JSON Web Token support on or off, depending on whether you want secure endpoints or not.
3. Run ```go install``` to install the required packages
4. Run ```go generate ./database/ent``` to update the ORM schema and to generate the corresponding files.


<b>Please note:</b> All files withing ```./database/ent``` except for ```./database/ent/schema``` are generated and should not be changed!

5. Compile and start the application by running ```go run --work main.go```

## Testing
Run ```go test ./tests -v``` to test the API. (Live server required)


