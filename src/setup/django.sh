docker-compose run --rm django sh -c "
  pip install django
  pip freeze > requirements.txt
"

echo
echo Create a project with the following command
echo docker-compose run --rm django django-admin startproject NAME .
echo
