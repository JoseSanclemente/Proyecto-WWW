from rest_framework import viewsets
from .models import Reads
from .serializer import ReadsSerializer


class ReadsViewSets(viewsets.ModelViewSet):
    queryset = Reads.objects.all()
    serializer_class = ReadsSerializer
