import os
from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from fastapi import File,UploadFile



def init():
    app = FastAPI()

    @app.get("/")
    def read_root():
        #get env variable
        return {"Hello": "World"}
    
    app.mount("/images",StaticFiles(directory="./img"),name="images")

    
    @app.post("/upload")
    async def upload_image(file: UploadFile = File(...)):
        #upload image
        current_directory = os.path.dirname(__file__)
        parent_directory = os.path.dirname(current_directory)

        uploaded_file = file.file.read()
        output_folder = os.path.join(parent_directory, "Storage/img")
        if not os.path.exists(output_folder):
            os.makedirs(output_folder)
        
        time_format = '%Y%m%d-%H%M%S'
        output_filename = f"{file.filename}"
        output_path = os.path.join(output_folder,output_filename)

        with open(output_path, 'wb') as output_file:
            output_file.write(uploaded_file)
        return {"message": "Image uploaded successfully", "filename": output_filename}


    return app

app = init()

if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=8000)