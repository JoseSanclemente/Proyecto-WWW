from rest_framework import serializers
from .models import Meter


# This class is used to transport this model through HTTP
class MeterSerializer(serializers.ModelSerializer):
    class Meta:
        model = Meter
        fields = '__all__'
