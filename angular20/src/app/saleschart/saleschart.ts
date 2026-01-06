import { Component, OnInit } from '@angular/core';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import { Productservice } from '../services/productservice';

@Component({
  selector: 'app-saleschart',
  templateUrl: './saleschart.html',
})
export class Saleschart implements OnInit {
  public chartSafeUrl: any;

  constructor(
    private productsService: Productservice,
    private sanitizer: DomSanitizer
  ) {}

  ngOnInit() {
    this.displaySalesChart();
  }

  displaySalesChart() {
    this.productsService.showSalesGraph().subscribe((htmlContent: string) => {
      // Standardize the content for a Data URL
      const encodedHtml = btoa(unescape(encodeURIComponent(htmlContent)));
      this.chartSafeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(
        `data:text/html;base64,${encodedHtml}`
      );
    });
  }
  
  // displaySalesChart() {
  //   this.productsService.showSalesGraph().subscribe((htmlContent: string) => {
  //     const blob = new Blob([htmlContent], { type: 'text/html' });
  //     this.chartSafeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(URL.createObjectURL(blob));
  //   });
  // }
}
