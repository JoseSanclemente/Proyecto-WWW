from rest_framework import serializers
from .models import Reads


# This class is used to transport this model through HTTP
class ReadsSerializer(serializers.ModelSerializer):
    class Meta:
        model = Reads
        fields = '__all__'
