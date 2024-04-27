FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip install --upgrade pip

RUN pip install --default-timeout=100 -r requirements.txt

COPY . .

CMD ["python", "main.py"]

EXPOSE 8000
