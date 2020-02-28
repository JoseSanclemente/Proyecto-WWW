from rest_framework import routers

from .viewsets import TransformerViewSets

router = routers.SimpleRouter()
router.register('Transformer', TransformerViewSets)

urlpatterns = router.urls
