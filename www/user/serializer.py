from rest_framework import serializers
from .models import User


# This class is used to transport this model through HTTP
class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = '__all__'
