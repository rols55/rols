# Use the official Node.js image as base
FROM node:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY /package*.json ./

# Install dependencies
RUN npm install

# Copy the source code from the current directory to the Working Directory inside the container
COPY / ./

RUN npm run build

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the application
CMD ["npm", "run", "start"]
