docker-compose run --rm rails sh -c "
  gem i rails
  rails new . -d postgresql -S
"
