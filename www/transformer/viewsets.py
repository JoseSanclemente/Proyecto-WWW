from rest_framework import viewsets
from .models import Transformer
from .serializer import TransformerSerializer


class TransformerViewSets(viewsets.ModelViewSet):
    queryset = Transformer.objects.all()
    serializer_class = TransformerSerializer
