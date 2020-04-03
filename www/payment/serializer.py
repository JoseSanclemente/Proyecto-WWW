from rest_framework import serializers
from .models import Payment


# This class is used to transport this model through HTTP
class PaymentSerializer(serializers.ModelSerializer):
    class Meta:
        model = Payment
        fields = '__all__'
