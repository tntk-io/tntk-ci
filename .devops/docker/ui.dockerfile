FROM node:12 as runtime
COPY . .
RUN npm install
CMD ["npm", "start"]
