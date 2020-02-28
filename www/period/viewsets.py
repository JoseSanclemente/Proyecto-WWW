from rest_framework import viewsets
from .models import Period
from .serializer import PeriodSerializer


class PeriodViewSets(viewsets.ModelViewSet):
    queryset = Period.objects.all()
    serializer_class = PeriodSerializer
