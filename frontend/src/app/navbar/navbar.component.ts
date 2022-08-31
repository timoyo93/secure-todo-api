import { Component, Input, OnInit } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss'],
})
export class NavbarComponent implements OnInit {
  @Input() logoTitle = 'LOGO';
  isAuthenticated = false;

  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit() {
    this.authService.isAuthenticated.subscribe(
      (v) => (this.isAuthenticated = v)
    );
  }

  logoutHandler() {
    this.authService.logoutUser().subscribe({
      next: (res) => {
        if (res.status === 200) {
          this.isAuthenticated = false;
          this.authService.isAuthenticated.next(false);
          this.router.navigate(['/', 'login']).then();
        }
      },
      error: (err) => console.log(err),
    });
  }
}
