from rest_framework import serializers
from .models import Electricaribe


# This class is used to transport this model through HTTP
class ElectricaribeSerializer(serializers.ModelSerializer):
    class Meta:
        model = Electricaribe
        fields = '__all__'
