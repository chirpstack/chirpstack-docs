build:
	docker-compose run --rm chirpstack-docs mdbook clean
	docker-compose run --rm chirpstack-docs mdbook build

serve:
	docker-compose run --rm --service-ports chirpstack-docs mdbook serve -n 0.0.0.0

upload:
	cd book && rsync -avzh . root@chirpstack.io:/var/www/chirpstack.io/docs/
