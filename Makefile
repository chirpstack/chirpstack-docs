build:
	mdbook build
	chmod 755 book/d2
	chmod 644 book/d2/*

dependencies:
	cargo install mdbook-findrep --locked --root .cargo
