from rest_framework import serializers
from .models import Metmodel


# This class is used to transport this model through HTTP
class MetmodelSerializer(serializers.ModelSerializer):
    class Meta:
        model = Metmodel
        fields = '__all__'
