docker-compose run --rm django sh -c "
  pip install django
  pip freeze > requirements.txt
"
