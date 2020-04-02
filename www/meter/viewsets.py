from rest_framework import viewsets
from .models import Meter
from .serializer import MeterSerializer


class MeterViewSets(viewsets.ModelViewSet):
    queryset = Meter.objects.all()
    serializer_class = MeterSerializer
