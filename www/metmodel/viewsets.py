from rest_framework import viewsets
from .models import Metmodel
from .serializer import MetmodelSerializer


class MetmodelViewSets(viewsets.ModelViewSet):
    queryset = Metmodel.objects.all()
    serializer_class = MetmodelSerializer
