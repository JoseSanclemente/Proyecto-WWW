from rest_framework import serializers
from .models import Transformer


# This class is used to transport this model through HTTP
class TransformerSerializer(serializers.ModelSerializer):
    class Meta:
        model = Transformer
        fields = '__all__'
