# Step 1: Build stage
FROM node:22 AS build

# Set working directory inside the container
WORKDIR /app

# Install dependencies and build the app
COPY package.json package-lock.json /app/
RUN npm install

# Copy source files and build the React app
COPY . /app/
RUN npm run build

# Step 2: Serve stage (using Nginx)
FROM nginx:alpine

# Copy the build files from the previous step
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]