from rest_framework import viewsets
from .models import Invoice
from .serializer import InvoiceSerializer


class InvoiceViewSets(viewsets.ModelViewSet):
    queryset = Invoice.objects.all()
    serializer_class = InvoiceSerializer
