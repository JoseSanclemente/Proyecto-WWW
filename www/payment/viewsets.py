from rest_framework import viewsets
from .models import Payment
from .serializer import PaymentSerializer


class PaymentViewSets(viewsets.ModelViewSet):
    queryset = Payment.objects.all()
    serializer_class = PaymentSerializer
