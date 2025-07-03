---
title: "Python Automation: From Repetitive Tasks to Production Workflows"
description: "Complete guide to Python automation covering file processing, API interactions, system administration, and building robust automation pipelines"
date: 2024-11-25
readingTime: 12
tags: ["Python", "Automation", "DevOps", "Scripting", "Workflows"]
author: "Billy Petersson"
featured: true
---

# Python Automation: From Repetitive Tasks to Production Workflows

Python's simplicity and extensive library ecosystem make it an excellent choice for automation. This comprehensive guide will take you from basic task automation to building sophisticated production workflows.

## Why Python for Automation?

Python excels at automation because of:

- **Readable syntax** that makes scripts easy to maintain
- **Rich standard library** covering most common tasks
- **Extensive third-party packages** for specialized operations
- **Cross-platform compatibility** for diverse environments
- **Strong community support** and documentation

## 1. File and Directory Operations

### Basic File Processing

Let's start with common file operations that form the foundation of many automation tasks:

```python
import os
import shutil
import glob
from pathlib import Path
import json
import csv

def organize_downloads_folder():
    """Organize files in downloads folder by extension."""
    downloads_path = Path.home() / "Downloads"
    
    # Define file type mappings
    file_types = {
        'images': ['.jpg', '.jpeg', '.png', '.gif', '.bmp', '.svg'],
        'documents': ['.pdf', '.doc', '.docx', '.txt', '.rtf'],
        'videos': ['.mp4', '.avi', '.mkv', '.mov', '.wmv'],
        'audio': ['.mp3', '.wav', '.flac', '.aac'],
        'archives': ['.zip', '.rar', '.tar', '.gz', '.7z'],
        'code': ['.py', '.js', '.html', '.css', '.java', '.cpp']
    }
    
    # Create directories if they don't exist
    for folder in file_types.keys():
        folder_path = downloads_path / folder
        folder_path.mkdir(exist_ok=True)
    
    # Move files to appropriate folders
    for file_path in downloads_path.iterdir():
        if file_path.is_file():
            file_extension = file_path.suffix.lower()
            
            moved = False
            for folder, extensions in file_types.items():
                if file_extension in extensions:
                    destination = downloads_path / folder / file_path.name
                    
                    # Handle duplicate names
                    counter = 1
                    while destination.exists():
                        name = file_path.stem
                        extension = file_path.suffix
                        destination = downloads_path / folder / f"{name}_{counter}{extension}"
                        counter += 1
                    
                    file_path.rename(destination)
                    print(f"Moved {file_path.name} to {folder}/")
                    moved = True
                    break
            
            if not moved:
                print(f"Unknown file type: {file_path.name}")

# Advanced file processing with metadata
def process_image_batch():
    """Batch process images with metadata extraction."""
    from PIL import Image, ExifTags
    import datetime
    
    image_dir = Path("./images")
    processed_dir = Path("./processed_images")
    processed_dir.mkdir(exist_ok=True)
    
    metadata_log = []
    
    for image_path in image_dir.glob("*.{jpg,jpeg,png}"):
        try:
            with Image.open(image_path) as img:
                # Extract EXIF data
                exif_data = {}
                if hasattr(img, '_getexif') and img._getexif():
                    exif = img._getexif()
                    for tag_id, value in exif.items():
                        tag = ExifTags.TAGS.get(tag_id, tag_id)
                        exif_data[tag] = value
                
                # Resize image
                max_size = (1920, 1080)
                img.thumbnail(max_size, Image.Lanczos)
                
                # Save processed image
                output_path = processed_dir / f"processed_{image_path.name}"
                img.save(output_path, optimize=True, quality=85)
                
                # Log metadata
                metadata_log.append({
                    'original_file': image_path.name,
                    'processed_file': output_path.name,
                    'original_size': image_path.stat().st_size,
                    'processed_size': output_path.stat().st_size,
                    'dimensions': img.size,
                    'timestamp': datetime.datetime.now().isoformat(),
                    'camera_make': exif_data.get('Make', 'Unknown'),
                    'camera_model': exif_data.get('Model', 'Unknown')
                })
                
                print(f"Processed {image_path.name}")
                
        except Exception as e:
            print(f"Error processing {image_path.name}: {e}")
    
    # Save metadata log
    with open(processed_dir / "processing_log.json", 'w') as f:
        json.dump(metadata_log, f, indent=2)

if __name__ == "__main__":
    organize_downloads_folder()
    process_image_batch()
```

## 2. Web Scraping and API Automation

### Advanced Web Scraping

```python
import requests
from bs4 import BeautifulSoup
import pandas as pd
import time
import logging
from dataclasses import dataclass
from typing import List, Optional
import concurrent.futures
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry

@dataclass
class Product:
    name: str
    price: float
    rating: Optional[float]
    reviews_count: int
    availability: str
    url: str

class EcommerceScraper:
    def __init__(self, base_url: str, headers: dict = None):
        self.base_url = base_url
        self.session = requests.Session()
        
        # Configure retry strategy
        retry_strategy = Retry(
            total=3,
            backoff_factor=1,
            status_forcelist=[429, 500, 502, 503, 504]
        )
        adapter = HTTPAdapter(max_retries=retry_strategy)
        self.session.mount("http://", adapter)
        self.session.mount("https://", adapter)
        
        # Set headers
        default_headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36'
        }
        self.session.headers.update(headers or default_headers)
        
        # Setup logging
        logging.basicConfig(level=logging.INFO)
        self.logger = logging.getLogger(__name__)
    
    def scrape_product_page(self, product_url: str) -> Optional[Product]:
        """Scrape individual product page."""
        try:
            response = self.session.get(product_url, timeout=10)
            response.raise_for_status()
            
            soup = BeautifulSoup(response.content, 'html.parser')
            
            # Extract product information (selectors depend on the website)
            name = soup.find('h1', class_='product-title')
            price = soup.find('span', class_='price')
            rating = soup.find('div', class_='rating')
            reviews = soup.find('span', class_='reviews-count')
            availability = soup.find('span', class_='availability')
            
            if not all([name, price]):
                return None
            
            return Product(
                name=name.text.strip(),
                price=float(price.text.replace('$', '').replace(',', '')),
                rating=float(rating.get('data-rating', 0)) if rating else None,
                reviews_count=int(reviews.text.split()[0]) if reviews else 0,
                availability=availability.text.strip() if availability else 'Unknown',
                url=product_url
            )
            
        except Exception as e:
            self.logger.error(f"Error scraping {product_url}: {e}")
            return None
    
    def scrape_category(self, category_url: str, max_pages: int = 5) -> List[Product]:
        """Scrape all products from a category."""
        products = []
        
        for page in range(1, max_pages + 1):
            try:
                page_url = f"{category_url}?page={page}"
                response = self.session.get(page_url, timeout=10)
                response.raise_for_status()
                
                soup = BeautifulSoup(response.content, 'html.parser')
                product_links = soup.find_all('a', class_='product-link')
                
                if not product_links:
                    break
                
                # Use threading for concurrent scraping
                with concurrent.futures.ThreadPoolExecutor(max_workers=5) as executor:
                    futures = []
                    for link in product_links:
                        product_url = self.base_url + link.get('href')
                        futures.append(executor.submit(self.scrape_product_page, product_url))
                    
                    for future in concurrent.futures.as_completed(futures):
                        product = future.result()
                        if product:
                            products.append(product)
                
                self.logger.info(f"Scraped page {page}, found {len(product_links)} products")
                time.sleep(1)  # Be respectful to the server
                
            except Exception as e:
                self.logger.error(f"Error scraping page {page}: {e}")
                break
        
        return products
    
    def save_to_csv(self, products: List[Product], filename: str):
        """Save products to CSV file."""
        df = pd.DataFrame([product.__dict__ for product in products])
        df.to_csv(filename, index=False)
        self.logger.info(f"Saved {len(products)} products to {filename}")

# API Integration Example
class APIAutomation:
    def __init__(self, api_key: str, base_url: str):
        self.api_key = api_key
        self.base_url = base_url
        self.session = requests.Session()
        self.session.headers.update({
            'Authorization': f'Bearer {api_key}',
            'Content-Type': 'application/json'
        })
    
    def bulk_user_operations(self, users_csv: str):
        """Perform bulk operations on users from CSV."""
        df = pd.read_csv(users_csv)
        results = []
        
        for _, user_data in df.iterrows():
            try:
                # Create user
                user_payload = {
                    'email': user_data['email'],
                    'name': user_data['name'],
                    'department': user_data['department']
                }
                
                response = self.session.post(
                    f"{self.base_url}/users",
                    json=user_payload
                )
                
                if response.status_code == 201:
                    user_id = response.json()['id']
                    
                    # Assign role
                    role_payload = {
                        'user_id': user_id,
                        'role': user_data['role']
                    }
                    
                    role_response = self.session.post(
                        f"{self.base_url}/roles/assign",
                        json=role_payload
                    )
                    
                    results.append({
                        'email': user_data['email'],
                        'user_created': True,
                        'role_assigned': role_response.status_code == 200,
                        'user_id': user_id
                    })
                else:
                    results.append({
                        'email': user_data['email'],
                        'user_created': False,
                        'error': response.text
                    })
                    
                time.sleep(0.1)  # Rate limiting
                
            except Exception as e:
                results.append({
                    'email': user_data['email'],
                    'user_created': False,
                    'error': str(e)
                })
        
        # Save results
        results_df = pd.DataFrame(results)
        results_df.to_csv('user_creation_results.csv', index=False)
        
        return results
```

## 3. System Administration Automation

### Server Monitoring and Maintenance

```python
import psutil
import subprocess
import smtplib
from email.mime.text import MimeText
from email.mime.multipart import MimeMultipart
import schedule
import time
from dataclasses import dataclass
from typing import Dict, List
import json
import datetime

@dataclass
class SystemMetrics:
    timestamp: str
    cpu_percent: float
    memory_percent: float
    disk_usage: Dict[str, float]
    network_io: Dict[str, int]
    running_processes: int
    load_average: List[float]

class SystemMonitor:
    def __init__(self, config_file: str = 'monitor_config.json'):
        with open(config_file, 'r') as f:
            self.config = json.load(f)
        
        self.thresholds = self.config['thresholds']
        self.email_config = self.config['email']
        self.metrics_history = []
    
    def collect_metrics(self) -> SystemMetrics:
        """Collect current system metrics."""
        # CPU usage
        cpu_percent = psutil.cpu_percent(interval=1)
        
        # Memory usage
        memory = psutil.virtual_memory()
        memory_percent = memory.percent
        
        # Disk usage for all mounted drives
        disk_usage = {}
        for partition in psutil.disk_partitions():
            try:
                usage = psutil.disk_usage(partition.mountpoint)
                disk_usage[partition.mountpoint] = usage.percent
            except PermissionError:
                continue
        
        # Network I/O
        network = psutil.net_io_counters()
        network_io = {
            'bytes_sent': network.bytes_sent,
            'bytes_recv': network.bytes_recv
        }
        
        # Running processes
        running_processes = len(psutil.pids())
        
        # Load average (Unix systems)
        try:
            load_average = list(psutil.getloadavg())
        except AttributeError:
            load_average = [0.0, 0.0, 0.0]
        
        return SystemMetrics(
            timestamp=datetime.datetime.now().isoformat(),
            cpu_percent=cpu_percent,
            memory_percent=memory_percent,
            disk_usage=disk_usage,
            network_io=network_io,
            running_processes=running_processes,
            load_average=load_average
        )
    
    def check_thresholds(self, metrics: SystemMetrics) -> List[str]:
        """Check if any metrics exceed thresholds."""
        alerts = []
        
        if metrics.cpu_percent > self.thresholds['cpu']:
            alerts.append(f"CPU usage: {metrics.cpu_percent:.1f}%")
        
        if metrics.memory_percent > self.thresholds['memory']:
            alerts.append(f"Memory usage: {metrics.memory_percent:.1f}%")
        
        for mount, usage in metrics.disk_usage.items():
            if usage > self.thresholds['disk']:
                alerts.append(f"Disk usage {mount}: {usage:.1f}%")
        
        if metrics.load_average[0] > self.thresholds['load']:
            alerts.append(f"Load average: {metrics.load_average[0]:.2f}")
        
        return alerts
    
    def send_alert_email(self, alerts: List[str], metrics: SystemMetrics):
        """Send alert email when thresholds are exceeded."""
        msg = MimeMultipart()
        msg['From'] = self.email_config['from']
        msg['To'] = ', '.join(self.email_config['to'])
        msg['Subject'] = f"System Alert - {datetime.datetime.now().strftime('%Y-%m-%d %H:%M')}"
        
        body = f"""
        System Alert Triggered!
        
        Timestamp: {metrics.timestamp}
        
        Threshold Violations:
        {chr(10).join('• ' + alert for alert in alerts)}
        
        Current System Status:
        • CPU Usage: {metrics.cpu_percent:.1f}%
        • Memory Usage: {metrics.memory_percent:.1f}%
        • Running Processes: {metrics.running_processes}
        • Load Average: {', '.join(f'{load:.2f}' for load in metrics.load_average)}
        
        Disk Usage:
        {chr(10).join(f'• {mount}: {usage:.1f}%' for mount, usage in metrics.disk_usage.items())}
        """
        
        msg.attach(MimeText(body, 'plain'))
        
        try:
            server = smtplib.SMTP(self.email_config['smtp_server'], self.email_config['smtp_port'])
            server.starttls()
            server.login(self.email_config['username'], self.email_config['password'])
            server.send_message(msg)
            server.quit()
            print("Alert email sent successfully")
        except Exception as e:
            print(f"Failed to send alert email: {e}")
    
    def cleanup_logs(self, log_dir: str, retention_days: int = 30):
        """Clean up old log files."""
        log_path = Path(log_dir)
        cutoff_date = datetime.datetime.now() - datetime.timedelta(days=retention_days)
        
        cleaned_files = []
        for log_file in log_path.glob('*.log'):
            if datetime.datetime.fromtimestamp(log_file.stat().st_mtime) < cutoff_date:
                log_file.unlink()
                cleaned_files.append(log_file.name)
        
        if cleaned_files:
            print(f"Cleaned up {len(cleaned_files)} old log files")
        
        return cleaned_files
    
    def backup_databases(self):
        """Backup databases using configured commands."""
        backup_dir = Path(self.config['backup']['directory'])
        backup_dir.mkdir(exist_ok=True)
        
        timestamp = datetime.datetime.now().strftime('%Y%m%d_%H%M%S')
        
        for db_config in self.config['backup']['databases']:
            try:
                backup_file = backup_dir / f"{db_config['name']}_{timestamp}.sql"
                
                cmd = [
                    'mysqldump',
                    '-h', db_config['host'],
                    '-u', db_config['user'],
                    f"-p{db_config['password']}",
                    db_config['database']
                ]
                
                with open(backup_file, 'w') as f:
                    subprocess.run(cmd, stdout=f, check=True)
                
                # Compress backup
                subprocess.run(['gzip', str(backup_file)], check=True)
                
                print(f"Database backup completed: {backup_file}.gz")
                
            except subprocess.CalledProcessError as e:
                print(f"Database backup failed for {db_config['name']}: {e}")
    
    def run_monitoring_cycle(self):
        """Run a complete monitoring cycle."""
        print(f"Running monitoring cycle at {datetime.datetime.now()}")
        
        # Collect metrics
        metrics = self.collect_metrics()
        self.metrics_history.append(metrics)
        
        # Keep only last 24 hours of metrics
        cutoff_time = datetime.datetime.now() - datetime.timedelta(hours=24)
        self.metrics_history = [
            m for m in self.metrics_history 
            if datetime.datetime.fromisoformat(m.timestamp) > cutoff_time
        ]
        
        # Check thresholds
        alerts = self.check_thresholds(metrics)
        if alerts:
            self.send_alert_email(alerts, metrics)
        
        # Save metrics to file
        with open('system_metrics.json', 'w') as f:
            json.dump([m.__dict__ for m in self.metrics_history], f, indent=2)
        
        print(f"Monitoring cycle completed. Alerts: {len(alerts)}")

# Configuration file example (monitor_config.json)
config_example = {
    "thresholds": {
        "cpu": 80,
        "memory": 85,
        "disk": 90,
        "load": 2.0
    },
    "email": {
        "smtp_server": "smtp.gmail.com",
        "smtp_port": 587,
        "username": "your_email@gmail.com",
        "password": "your_app_password",
        "from": "monitor@yourserver.com",
        "to": ["admin@yourcompany.com"]
    },
    "backup": {
        "directory": "/backups",
        "databases": [
            {
                "name": "production_db",
                "host": "localhost",
                "user": "backup_user",
                "password": "backup_password",
                "database": "production"
            }
        ]
    }
}
```

## 4. Data Processing Automation

### ETL Pipeline Implementation

```python
import pandas as pd
import sqlite3
from sqlalchemy import create_engine
import requests
import logging
from pathlib import Path
import json
from typing import Dict, List, Any
from dataclasses import dataclass
import boto3
from datetime import datetime, timedelta

@dataclass
class DataSource:
    name: str
    type: str  # 'api', 'csv', 'database', 's3'
    connection_params: Dict[str, Any]
    query: str = None

class ETLPipeline:
    def __init__(self, config_file: str):
        with open(config_file, 'r') as f:
            self.config = json.load(f)
        
        self.setup_logging()
        self.setup_connections()
    
    def setup_logging(self):
        """Setup logging configuration."""
        logging.basicConfig(
            level=logging.INFO,
            format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
            handlers=[
                logging.FileHandler('etl_pipeline.log'),
                logging.StreamHandler()
            ]
        )
        self.logger = logging.getLogger(__name__)
    
    def setup_connections(self):
        """Setup database and other service connections."""
        self.engines = {}
        
        for db_name, db_config in self.config['databases'].items():
            connection_string = (
                f"{db_config['type']}://{db_config['user']}:"
                f"{db_config['password']}@{db_config['host']}:"
                f"{db_config['port']}/{db_config['database']}"
            )
            self.engines[db_name] = create_engine(connection_string)
        
        # Setup AWS S3 client if configured
        if 'aws' in self.config:
            self.s3_client = boto3.client(
                's3',
                aws_access_key_id=self.config['aws']['access_key'],
                aws_secret_access_key=self.config['aws']['secret_key'],
                region_name=self.config['aws']['region']
            )
    
    def extract_from_api(self, source: DataSource) -> pd.DataFrame:
        """Extract data from API source."""
        try:
            headers = source.connection_params.get('headers', {})
            params = source.connection_params.get('params', {})
            
            response = requests.get(
                source.connection_params['url'],
                headers=headers,
                params=params,
                timeout=30
            )
            response.raise_for_status()
            
            data = response.json()
            
            # Handle nested JSON structures
            if 'data_path' in source.connection_params:
                for key in source.connection_params['data_path'].split('.'):
                    data = data[key]
            
            df = pd.json_normalize(data)
            self.logger.info(f"Extracted {len(df)} records from API: {source.name}")
            
            return df
            
        except Exception as e:
            self.logger.error(f"Failed to extract from API {source.name}: {e}")
            return pd.DataFrame()
    
    def extract_from_database(self, source: DataSource) -> pd.DataFrame:
        """Extract data from database source."""
        try:
            engine = self.engines[source.connection_params['engine_name']]
            df = pd.read_sql(source.query, engine)
            
            self.logger.info(f"Extracted {len(df)} records from database: {source.name}")
            return df
            
        except Exception as e:
            self.logger.error(f"Failed to extract from database {source.name}: {e}")
            return pd.DataFrame()
    
    def extract_from_csv(self, source: DataSource) -> pd.DataFrame:
        """Extract data from CSV file."""
        try:
            file_path = source.connection_params['file_path']
            
            # Handle S3 CSV files
            if file_path.startswith('s3://'):
                bucket, key = file_path[5:].split('/', 1)
                obj = self.s3_client.get_object(Bucket=bucket, Key=key)
                df = pd.read_csv(obj['Body'])
            else:
                df = pd.read_csv(file_path)
            
            self.logger.info(f"Extracted {len(df)} records from CSV: {source.name}")
            return df
            
        except Exception as e:
            self.logger.error(f"Failed to extract from CSV {source.name}: {e}")
            return pd.DataFrame()
    
    def transform_data(self, df: pd.DataFrame, transformations: List[Dict]) -> pd.DataFrame:
        """Apply transformations to dataframe."""
        for transform in transformations:
            transform_type = transform['type']
            
            try:
                if transform_type == 'rename_columns':
                    df = df.rename(columns=transform['mapping'])
                
                elif transform_type == 'filter_rows':
                    condition = transform['condition']
                    df = df.query(condition)
                
                elif transform_type == 'add_calculated_column':
                    df[transform['column_name']] = df.eval(transform['expression'])
                
                elif transform_type == 'convert_types':
                    for col, dtype in transform['columns'].items():
                        if dtype == 'datetime':
                            df[col] = pd.to_datetime(df[col])
                        else:
                            df[col] = df[col].astype(dtype)
                
                elif transform_type == 'fill_missing':
                    method = transform['method']
                    if method == 'forward_fill':
                        df = df.fillna(method='ffill')
                    elif method == 'backward_fill':
                        df = df.fillna(method='bfill')
                    elif method == 'value':
                        df = df.fillna(transform['value'])
                
                elif transform_type == 'aggregate':
                    group_by = transform['group_by']
                    agg_funcs = transform['aggregations']
                    df = df.groupby(group_by).agg(agg_funcs).reset_index()
                
                elif transform_type == 'join':
                    other_df = self.extract_data(transform['source'])
                    df = df.merge(
                        other_df,
                        on=transform['on'],
                        how=transform.get('how', 'inner')
                    )
                
                self.logger.info(f"Applied transformation: {transform_type}")
                
            except Exception as e:
                self.logger.error(f"Failed to apply transformation {transform_type}: {e}")
        
        return df
    
    def load_to_database(self, df: pd.DataFrame, table_name: str, engine_name: str, if_exists: str = 'replace'):
        """Load dataframe to database table."""
        try:
            engine = self.engines[engine_name]
            df.to_sql(table_name, engine, if_exists=if_exists, index=False)
            
            self.logger.info(f"Loaded {len(df)} records to table: {table_name}")
            
        except Exception as e:
            self.logger.error(f"Failed to load to database table {table_name}: {e}")
    
    def load_to_csv(self, df: pd.DataFrame, file_path: str):
        """Load dataframe to CSV file."""
        try:
            if file_path.startswith('s3://'):
                # Upload to S3
                bucket, key = file_path[5:].split('/', 1)
                csv_buffer = df.to_csv(index=False)
                self.s3_client.put_object(
                    Bucket=bucket,
                    Key=key,
                    Body=csv_buffer
                )
            else:
                df.to_csv(file_path, index=False)
            
            self.logger.info(f"Loaded {len(df)} records to CSV: {file_path}")
            
        except Exception as e:
            self.logger.error(f"Failed to load to CSV {file_path}: {e}")
    
    def run_pipeline(self, pipeline_name: str):
        """Run a complete ETL pipeline."""
        pipeline_config = self.config['pipelines'][pipeline_name]
        
        self.logger.info(f"Starting ETL pipeline: {pipeline_name}")
        
        # Extract data from all sources
        dataframes = {}
        for source_config in pipeline_config['sources']:
            source = DataSource(**source_config)
            
            if source.type == 'api':
                df = self.extract_from_api(source)
            elif source.type == 'database':
                df = self.extract_from_database(source)
            elif source.type == 'csv':
                df = self.extract_from_csv(source)
            else:
                self.logger.error(f"Unknown source type: {source.type}")
                continue
            
            dataframes[source.name] = df
        
        # Transform data
        main_df = dataframes[pipeline_config['main_source']]
        if 'transformations' in pipeline_config:
            main_df = self.transform_data(main_df, pipeline_config['transformations'])
        
        # Load data to destinations
        for destination in pipeline_config['destinations']:
            dest_type = destination['type']
            
            if dest_type == 'database':
                self.load_to_database(
                    main_df,
                    destination['table_name'],
                    destination['engine_name'],
                    destination.get('if_exists', 'replace')
                )
            elif dest_type == 'csv':
                self.load_to_csv(main_df, destination['file_path'])
        
        self.logger.info(f"Completed ETL pipeline: {pipeline_name}")
        
        return main_df

# Usage example
if __name__ == "__main__":
    # Setup monitoring
    monitor = SystemMonitor('monitor_config.json')
    
    # Schedule monitoring tasks
    schedule.every(5).minutes.do(monitor.run_monitoring_cycle)
    schedule.every().day.at("02:00").do(monitor.backup_databases)
    schedule.every().day.at("03:00").do(lambda: monitor.cleanup_logs('/var/log', 30))
    
    # Run ETL pipeline
    etl = ETLPipeline('etl_config.json')
    
    # Schedule ETL jobs
    schedule.every().hour.do(lambda: etl.run_pipeline('hourly_sales_data'))
    schedule.every().day.at("01:00").do(lambda: etl.run_pipeline('daily_user_analytics'))
    
    print("Automation system started. Press Ctrl+C to stop.")
    
    try:
        while True:
            schedule.run_pending()
            time.sleep(60)
    except KeyboardInterrupt:
        print("Automation system stopped.")
```

## 5. Production-Ready Automation

### Error Handling and Monitoring

```python
import functools
import traceback
from typing import Callable, Any
import time
import logging
from dataclasses import dataclass
from enum import Enum

class TaskStatus(Enum):
    PENDING = "pending"
    RUNNING = "running"
    SUCCESS = "success"
    FAILED = "failed"
    RETRY = "retry"

@dataclass
class TaskResult:
    task_name: str
    status: TaskStatus
    start_time: float
    end_time: float = None
    error_message: str = None
    retry_count: int = 0
    result: Any = None

def retry_on_failure(max_retries: int = 3, delay: float = 1.0, backoff: float = 2.0):
    """Decorator for automatic retry with exponential backoff."""
    def decorator(func: Callable) -> Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            last_exception = None
            
            for attempt in range(max_retries + 1):
                try:
                    return func(*args, **kwargs)
                except Exception as e:
                    last_exception = e
                    if attempt < max_retries:
                        wait_time = delay * (backoff ** attempt)
                        logging.warning(
                            f"Attempt {attempt + 1} failed for {func.__name__}: {e}. "
                            f"Retrying in {wait_time:.2f} seconds..."
                        )
                        time.sleep(wait_time)
                    else:
                        logging.error(
                            f"All {max_retries + 1} attempts failed for {func.__name__}: {e}"
                        )
            
            raise last_exception
        return wrapper
    return decorator

def monitor_task(task_name: str):
    """Decorator for task monitoring and logging."""
    def decorator(func: Callable) -> Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            task_result = TaskResult(
                task_name=task_name,
                status=TaskStatus.RUNNING,
                start_time=time.time()
            )
            
            logger = logging.getLogger(__name__)
            logger.info(f"Starting task: {task_name}")
            
            try:
                result = func(*args, **kwargs)
                task_result.status = TaskStatus.SUCCESS
                task_result.result = result
                task_result.end_time = time.time()
                
                duration = task_result.end_time - task_result.start_time
                logger.info(f"Task completed successfully: {task_name} ({duration:.2f}s)")
                
                return result
                
            except Exception as e:
                task_result.status = TaskStatus.FAILED
                task_result.error_message = str(e)
                task_result.end_time = time.time()
                
                duration = task_result.end_time - task_result.start_time
                logger.error(
                    f"Task failed: {task_name} ({duration:.2f}s)\n"
                    f"Error: {e}\n"
                    f"Traceback: {traceback.format_exc()}"
                )
                
                raise
            
        return wrapper
    return decorator

# Example usage of decorators
@monitor_task("data_backup")
@retry_on_failure(max_retries=3, delay=2.0)
def backup_critical_data():
    """Example task with monitoring and retry logic."""
    # Simulate potential failure
    import random
    if random.random() < 0.3:  # 30% chance of failure
        raise Exception("Simulated backup failure")
    
    print("Backup completed successfully")
    return "backup_file_20241225.tar.gz"

if __name__ == "__main__":
    # Example usage
    try:
        result = backup_critical_data()
        print(f"Backup result: {result}")
    except Exception as e:
        print(f"Backup ultimately failed: {e}")
```

## Best Practices for Production Automation

### 1. Configuration Management

```python
# Use environment variables and config files
import os
from pathlib import Path

class Config:
    def __init__(self, config_file: str = None):
        self.config_file = config_file or os.getenv('AUTOMATION_CONFIG', 'config.json')
        self.load_config()
    
    def load_config(self):
        """Load configuration from file and environment variables."""
        if Path(self.config_file).exists():
            with open(self.config_file, 'r') as f:
                self.config = json.load(f)
        else:
            self.config = {}
        
        # Override with environment variables
        self.config.update({
            'database_url': os.getenv('DATABASE_URL', self.config.get('database_url')),
            'api_key': os.getenv('API_KEY', self.config.get('api_key')),
            'log_level': os.getenv('LOG_LEVEL', self.config.get('log_level', 'INFO'))
        })
```

### 2. Logging and Monitoring

```python
import logging
import json
from datetime import datetime

class StructuredLogger:
    def __init__(self, name: str):
        self.logger = logging.getLogger(name)
        handler = logging.StreamHandler()
        handler.setFormatter(logging.Formatter('%(message)s'))
        self.logger.addHandler(handler)
        self.logger.setLevel(logging.INFO)
    
    def log_structured(self, level: str, message: str, **kwargs):
        """Log structured data in JSON format."""
        log_entry = {
            'timestamp': datetime.utcnow().isoformat(),
            'level': level,
            'message': message,
            **kwargs
        }
        
        getattr(self.logger, level.lower())(json.dumps(log_entry))
```

### 3. Testing Automation Scripts

```python
import unittest
from unittest.mock import patch, MagicMock
import tempfile
import os

class TestAutomationScripts(unittest.TestCase):
    def setUp(self):
        self.temp_dir = tempfile.mkdtemp()
    
    def tearDown(self):
        import shutil
        shutil.rmtree(self.temp_dir)
    
    @patch('requests.get')
    def test_api_extraction(self, mock_get):
        """Test API data extraction."""
        # Mock API response
        mock_response = MagicMock()
        mock_response.json.return_value = {'data': [{'id': 1, 'name': 'test'}]}
        mock_response.status_code = 200
        mock_get.return_value = mock_response
        
        # Test extraction
        source = DataSource(
            name='test_api',
            type='api',
            connection_params={'url': 'https://api.example.com/data'}
        )
        
        etl = ETLPipeline('test_config.json')
        df = etl.extract_from_api(source)
        
        self.assertEqual(len(df), 1)
        self.assertEqual(df.iloc[0]['name'], 'test')
    
    def test_file_organization(self):
        """Test file organization function."""
        # Create test files
        test_files = ['test.jpg', 'document.pdf', 'video.mp4']
        for filename in test_files:
            (Path(self.temp_dir) / filename).touch()
        
        # Test organization (would need to adapt organize_downloads_folder)
        # organize_files(self.temp_dir)
        
        # Assert files are moved to correct directories
        self.assertTrue((Path(self.temp_dir) / 'images' / 'test.jpg').exists())
        self.assertTrue((Path(self.temp_dir) / 'documents' / 'document.pdf').exists())

if __name__ == '__main__':
    unittest.main()
```

## Conclusion

Python automation can transform repetitive tasks into reliable, scalable processes. Key takeaways:

1. **Start Simple**: Begin with basic file operations and gradually add complexity
2. **Error Handling**: Always implement proper error handling and retry logic
3. **Logging**: Comprehensive logging is crucial for debugging and monitoring
4. **Configuration**: Use external configuration files and environment variables
5. **Testing**: Write tests for your automation scripts
6. **Monitoring**: Implement monitoring and alerting for production systems

Remember that good automation scripts are:
- **Reliable**: Handle errors gracefully and recover when possible
- **Maintainable**: Well-documented and easy to modify
- **Secure**: Protect sensitive information and validate inputs
- **Efficient**: Optimize for performance and resource usage

---

*What automation tasks are you working on? Share your experiences and challenges in the comments below!*

## Further Reading

- [Python Official Documentation](https://docs.python.org/3/)
- [Pandas Documentation](https://pandas.pydata.org/docs/)
- [Requests Library](https://requests.readthedocs.io/)
- [Schedule Library](https://schedule.readthedocs.io/)