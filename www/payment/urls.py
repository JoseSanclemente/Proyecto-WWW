from rest_framework import routers

from .viewsets import PaymentViewSets

router = routers.SimpleRouter()
router.register('Payment', PaymentViewSets)

urlpatterns = router.urls
