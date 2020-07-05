docker-compose run --rm rails bash -c "
  gem i rails
  rails new . -d postgresql -S
"
