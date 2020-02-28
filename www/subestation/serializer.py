from rest_framework import serializers
from .models import Subestation


# This class is used to transport this model through HTTP
class SubestationSerializer(serializers.ModelSerializer):
    class Meta:
        model = Subestation
        fields = '__all__'