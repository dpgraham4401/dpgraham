from django.db import models

# Create your models here.

STATUS = (
    (0, "Draft"),
    (1, "Publish")
)

class BlogPost(models.Model):
    title = models.CharField(max_length=255, unique=True)
    slug = models.SlugField(max_length=255, unique=True)
    created_date = models.DateTimeField(auto_now_add=True)
    last_update_date = models.DateField()
    content = models.TextField()
    author = models.CharField(max_length=255, default="David Graham")
    status = models.IntegerField(choices=STATUS, default=0)

    class Meta:
        ordering = ['-created_date']

    def __str__(self):
        return self.title
