import { Component, OnInit, ChangeDetectorRef, signal } from '@angular/core';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import { Productservice } from '../services/productservice';
import { MatProgressBarModule } from '@angular/material/progress-bar';

@Component({
  selector: 'app-saleschart',
  standalone: true,
  imports: [MatProgressBarModule],  
  templateUrl: './saleschart.html',
})
export class Saleschart implements OnInit {
  public chartSafeUrl: any;
  isLoading = signal(true); // Define as a signal

  constructor(
    private productsService: Productservice,
    private sanitizer: DomSanitizer,
    private cdr: ChangeDetectorRef // Inject this
  ) {}
  
  onIframeLoad() {
    this.isLoading.set(false);
    this.cdr.detectChanges(); // Force UI update
  }
  // constructor(
  //   private productsService: Productservice,
  //   private sanitizer: DomSanitizer
  // ) {}

  ngOnInit() {
    this.displaySalesChart();
  }

  displaySalesChart() {
    this.productsService.showSalesGraph().subscribe((htmlContent: string) => {
      const encodedHtml = btoa(unescape(encodeURIComponent(htmlContent)));
      this.chartSafeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(
        `data:text/html;base64,${encodedHtml}`
      );
    });
  }
  
}
