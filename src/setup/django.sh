docker-compose run --rm django bash -c "
  pip install django
  pip freeze > requirements.txt
"
