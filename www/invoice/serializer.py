from rest_framework import serializers
from .models import Invoice


# This class is used to transport this model through HTTP
class InvoiceSerializer(serializers.ModelSerializer):
    class Meta:
        model = Invoice
        fields = '__all__'
