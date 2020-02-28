from rest_framework import serializers
from .models import Contract


# This class is used to transport this model through HTTP
class ContractSerializer(serializers.ModelSerializer):
    class Meta:
        model = Contract
        fields = '__all__'
