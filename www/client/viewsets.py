from rest_framework import viewsets
from .models import Client
from .serializer import ClientSerializer


class ClientViewSets(viewsets.ModelViewSet):
    queryset = Client.objects.all()
    serializer_class = ClientSerializer
