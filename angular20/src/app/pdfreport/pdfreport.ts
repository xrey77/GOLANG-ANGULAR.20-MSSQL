import { Component, OnInit, inject, DestroyRef } from '@angular/core';
import { Productservice } from '../services/productservice';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';

@Component({
  selector: 'app-pdfreport',
  standalone: true, 
  imports: [], // Add CommonModule or others if needed for the template
  templateUrl: './pdfreport.html',
  styleUrl: './pdfreport.scss',
})
export class Pdfreport implements OnInit {
  // Use inject() for modern DI
  private productsService = inject(Productservice);
  private sanitizer = inject(DomSanitizer);
  private destroyRef = inject(DestroyRef);

  pdfUrl: SafeResourceUrl | null = null;
  private currentBlobUrl: string | null = null;

  ngOnInit() { // Corrected from gOnInit
    this.onViewReport();

    // Ensure memory is cleared when component is destroyed
    this.destroyRef.onDestroy(() => {
      if (this.currentBlobUrl) {
        URL.revokeObjectURL(this.currentBlobUrl);
      }
    });
  }

  onViewReport() {
    this.productsService.showPdfReport().subscribe({
      next: (blob: Blob) => {
        // Clean up previous blob to prevent memory bloat
        if (this.currentBlobUrl) {
          URL.revokeObjectURL(this.currentBlobUrl);
        }

        this.currentBlobUrl = URL.createObjectURL(blob);
        this.pdfUrl = this.sanitizer.bypassSecurityTrustResourceUrl(this.currentBlobUrl);
      },
      error: (err) => console.error('Error downloading PDF:', err)
    });
  }
}
