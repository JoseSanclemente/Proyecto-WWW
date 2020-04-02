from rest_framework import serializers
from .models import Employee


# This class is used to transport this model through HTTP
class EmployeeSerializer(serializers.ModelSerializer):
    class Meta:
        model = Employee
        fields = '__all__'

        
