from rest_framework import serializers
from .models import Client


# This class is used to transport this model through HTTP
class ClientSerializer(serializers.ModelSerializer):
    class Meta:
        model = Client
        fields = '__all__'
