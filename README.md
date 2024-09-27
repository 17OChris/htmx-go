# Setup
clone repo `git clone git@github.com:17OChris/htmx-go.git`

# Build and run with Docker

```bash
# Build the image
docker build -t htmx .
# Run a container from the image we just created
run `docker run -p 80:80 htmx`
```


# Tailwind dev setup
If you want to play around with the styles locally - run the below command.

This command will install tailwind and watch file for hotreload.

```bash
yarn

make css-watch
```