import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.enableCors();

  const config = new DocumentBuilder()
    .setTitle('API Gateway Service')
    .setDescription('API Gateway port that proxes requests to other microservices')
    .setVersion('1.0')
    .addTag('users - users service')
    .addTag('products - products service')
    .addTag('orders - orders service')
    .addTag('reviews - reviews service')
    .addTag('chat - chat service')
    .addTag('auth - auth/security service')
    .addBearerAuth()
    .build();

  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('api', app, document);

  await app.listen(process.env.PORT ?? 3000);
}
bootstrap();
