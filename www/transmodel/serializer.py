from rest_framework import serializers
from .models import Transmodel


# This class is used to transport this model through HTTP
class TransmodelSerializer(serializers.ModelSerializer):
    class Meta:
        model = Transmodel
        fields = '__all__'
