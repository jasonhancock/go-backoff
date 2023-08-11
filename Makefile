all:
	true

# Renders the README from the template. Uses the becca tool from
# https://github.com/dave/rebecca
readme:
	becca -package github.com/jasonhancock/go-backoff
