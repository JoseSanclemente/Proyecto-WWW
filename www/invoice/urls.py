from rest_framework import routers

from .viewsets import InvoiceViewSets

router = routers.SimpleRouter()
router.register('invoice', InvoiceViewSets)

urlpatterns = router.urls
