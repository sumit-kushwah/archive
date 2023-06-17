# Update System
sudo apt update - y
sudo apt --upgrade -y 

# Python3 and Pip3 sanity checking

which python3 # if not installed install using sudo apt install python3 -y
which pip3 # if not installed install using sudo apt install python3-pip -y

which virtualenv # if not installed install using pip install virtualenv and virtualenv bin path to $PATH
# export PATH=~/.local/bin:$PATH # add this line to .bashrc file if virutalenv installed in local home bin folder

# create new project in ~/english-to-indic-fastapi
mkdir ~/english-to-indic-fastapi

# change directory 
cd  ~/english-to-indic-fastapi

# create virtualenv folder as venv
virtualenv venv

# activate virtual enviroment
. venv/bin/activate

# run below command to makesure you are using correct python3 and pip3
which python3
which pip3

# create requirements.txt, gunicorn_config.py, app.py

# pip install 
pip3 install fastapi uvicorn gunicorn

pip freeze > requirements.txt

# gunicorn_config.py file
# Gunicorn config params
bind = "0.0.0.0:5000"
workers = 2
timeout = 1200
proc_name = 'gunicorn.proc'
pidfile = '/tmp/gunicorn.pid'
logfile = '/var/logs/debug.log'
loglevel = 'debug'
errorlog = '/var/logs/error.log'
accesslog = '/var/logs/access.log'
max_requests = 10000  # kill worker after max requests
# DEBUG = True

# uvicorn-specific option
worker_class = "uvicorn.workers.UvicornWorker"

# app.py file content
from fastapi import FastAPI

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}
    
# create logs folder
sudo mkdir /var/logs
sudo touch /var/logs/access.log
sudo touch /var/logs/error.log
sudo touch /var/logs/debug.log

# run command to test if everything is okay
sudo /home/$USER/english-to-indic-fastapi/venv/bin/gunicorn app:app --config gunicorn_config.py 

# installing supervisor
sudo apt install supervisor

# add below line in /etc/supervisor/supervisord.conf file
[inet_http_server]  ; Add this section for web server access
port=9001  ; Port number for web server access (choose a desired port)
username=user  ; Username for web server authentication
password=pass  ; Password for web server authentication

# restart supervisor service
sudo systemctl restart supervisor
# check status
sudo systemctl status supervisor

# create supervisor gunicorn.conf file in conf.d folder with content
[program:english_to_indic_gunicorn]
user=root
directory=/home/umendrasingh028/english-to-indic-fastapi
command=/home/umendrasingh028/english-to-indic-fastapi/venv/bin/gunicorn app:app --config gunicorn_config.py
autostart=true
autorestart=true
stdout_logfile=/var/log/en2indic/gunicorn.log
stderr_logfile=/var/log/en2indic/gunicorn_err.log

# create indicTrans_setup.sh file with following content
# make sure install git before running this script
sudo apt install git -y

# run install_model.sh file to install model and unzip it.
# install unzip for downloading the model
sudo apt install unzip -y
