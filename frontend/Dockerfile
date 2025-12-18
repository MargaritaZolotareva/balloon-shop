FROM node:16-alpine

WORKDIR /app

RUN npm install -g @vue/cli

RUN vue create -d frontend

WORKDIR /app/frontend

RUN npm install

EXPOSE 8080

CMD ["npm", "run", "serve"]