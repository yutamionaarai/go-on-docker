FROM node:16.14.2-alpine
# Set the current working directory inside the container
WORKDIR /go/src/app/frontend/my-app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the source code
COPY . .

# Expose port 3000 to the outside world
EXPOSE 3000

# Run the app
CMD ["npm", "run", "dev"]