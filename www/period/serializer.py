from rest_framework import serializers
from .models import Period


# This class is used to transport this model through HTTP
class PeriodSerializer(serializers.ModelSerializer):
    class Meta:
        model = Period
        fields = '__all__'
