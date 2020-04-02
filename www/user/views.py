from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from captcha.image import ImageCaptcha
from .models import User
from .serializer import UserSerializer
import base64
import random

class UserAPIView(APIView):
    def post(self, request, format=None):
        #serializer = UserSerializer(data=request.data)
        
        #if not serializer.is_valid():
        #    return Response(status=status.HTTP_404_NOT_FOUND)
        
        username = request.data.get('username')
        password = request.data.get('password')

        try:
            user = User.objects.get(pk=username)
            if password == user.password:
                return Response({'message': 'ok'})

            return Response(status=status.HTTP_404_NOT_FOUND)
            
        except User.DoesNotExist:
            return Response(status=status.HTTP_404_NOT_FOUND)

        return Response(status=status.HTTP_500_INTERNAL_SERVER_ERROR)

    def put(self, request, format=None):
        serializer = UserSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class CaptchaAPIView(APIView):
    def get(self, request, format=None):
        image = ImageCaptcha()

        letters = 'abcdefghijklmnopqrstuvwxyz1234567890'
        captcha = ''

        for _ in range(6):
            captcha += random.choice(letters)

        data = image.generate(captcha)
        return Response({'captcha': captcha, 'image': base64.b64encode(data.read())})