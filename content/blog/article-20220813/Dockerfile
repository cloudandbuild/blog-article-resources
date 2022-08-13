FROM jupyter/datascience-notebook:r-4.1.3
WORKDIR /home/jovyan
COPY requirements.txt ./
RUN pip install -r requirements.txt
COPY app.py ./
COPY notebooks ./notebooks
CMD ["python", "app.py"]
