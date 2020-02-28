from rest_framework import viewsets
from .models import Subestation
from .serializer import SubestationSerializer


class SubestationViewSets(viewsets.ModelViewSet):
    queryset = Subestation.objects.all()
    serializer_class = SubestationSerializer
