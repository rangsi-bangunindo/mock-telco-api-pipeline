import os
from dotenv import load_dotenv

load_dotenv()  # Loads variables from .env into os.environ

pg_host = os.getenv("PG_HOST")
pg_user = os.getenv("PG_USER")
pg_password = os.getenv("PG_PASSWORD")
pg_database = os.getenv("PG_DATABASE")