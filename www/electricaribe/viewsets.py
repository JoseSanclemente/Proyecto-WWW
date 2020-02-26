from rest_framework import viewsets
from .models import Electricaribe
from .serializer import ElectricaribeSerializer


class ElectricaribeViewSets(viewsets.ModelViewSet):
    queryset = Electricaribe.objects.all()
    serializer_class = ElectricaribeSerializer
