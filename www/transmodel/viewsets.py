from rest_framework import viewsets
from .models import Transmodel
from .serializer import TransmodelSerializer


class TransmodelViewSets(viewsets.ModelViewSet):
    queryset = Transmodel.objects.all()
    serializer_class = TransmodelSerializer
